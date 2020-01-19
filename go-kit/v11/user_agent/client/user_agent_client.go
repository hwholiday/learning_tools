package main

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"learning_tools/go-kit/v11/user_agent/pb"
	"learning_tools/go-kit/v11/user_agent/src"
	"learning_tools/go-kit/v11/utils"
	"time"
)

type UserAgent struct {
	instancerm *etcdv3.Instancer
	logger     log.Logger
	tracer     opentracing.Tracer
}

func NewUserAgentClient(addr []string, logger log.Logger) (*UserAgent, error) {
	var (
		etcdAddrs = addr
		serName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	tracer, _, err := utils.NewJaegerTracer("user_agent_client")
	if err != nil {
		return nil, err
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		return nil, err
	}
	instancerm, err := etcdv3.NewInstancer(etcdClient, serName, logger)
	if err != nil {
		return nil, err
	}
	return &UserAgent{
		instancerm: instancerm,
		logger:     logger,
		tracer:     tracer,
	}, err
}

func (u *UserAgent) UserAgentClient() (src.Service, error) {
	var (
		retryMax     = 3
		retryTimeout = 5 * time.Second
	)
	var (
		endpoints src.EndPointServer
	)
	{
		factory := u.factoryFor(src.MakeLoginEndPoint)
		endpointer := sd.NewEndpointer(u.instancerm, factory, u.logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.LoginEndPoint = retry
	}
	return endpoints, nil
}

func (u *UserAgent) factoryFor(makeEndpoint func(src.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		chainUnaryServer := grpcmiddleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(u.tracer)),
			grpc_zap.UnaryClientInterceptor(utils.GetLogger()),
			//utils.JaegerServerMiddleware(tracer),
		)
		conn, err := grpc.Dial(instance, grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(chainUnaryServer))
		/*utils.JaegerClientMiddleware(u.tracer)),*/
		if err != nil {
			return nil, nil, err
		}
		srv := u.NewGRPCClient(conn)
		endpoints := makeEndpoint(srv)
		return endpoints, conn, err
	}
}

func (u *UserAgent) NewGRPCClient(conn *grpc.ClientConn) src.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
			md.Set(utils.ContextReqUUid, UUID)
			ctx = metadata.NewOutgoingContext(ctx, *md)
			return ctx
		}),
	}
	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"pb.User",
			"RpcUserLogin",
			u.RequestLogin,
			u.ResponseLogin,
			pb.LoginAck{},
			options...).Endpoint()
	}
	return src.EndPointServer{
		LoginEndPoint: loginEndpoint,
	}
}

func (u *UserAgent) RequestLogin(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(src.LoginRequest)
	return &pb.Login{Account: req.In.Account, Password: req.In.Password}, nil
}

func (u *UserAgent) ResponseLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginAck)
	res := src.LoginResponse{}
	if resp.Err != "" {
		res.Err = errors.New(resp.Err)
	} else {
		res.Err = nil
		res.Ack = src.LoginAck{Token: resp.Token}
	}
	return res, nil
}
