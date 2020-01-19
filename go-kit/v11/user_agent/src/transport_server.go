package src

import (
	"context"
	"fmt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"learning_tools/go-kit/v11/user_agent/pb"
	"learning_tools/go-kit/v11/utils"
)

type grpcServer struct {
	login grpctransport.Handler
}

func NewGRPCServer(endpoint EndPointServer, log *zap.Logger) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, utils.ContextReqUUid, md.Get(utils.ContextReqUUid))
			return ctx
		}),
		grpctransport.ServerErrorHandler(utils.NewZapLogErrorHandler(log)),
	}
	return &grpcServer{login: grpctransport.NewServer(
		endpoint.LoginEndPoint,
		RequestGrpcLogin,
		ResponseGrpcLogin,
		options...,
	)}
}

func (s *grpcServer) RpcUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		fmt.Println("s.login.ServeGRPC", err)
		return nil, err
	}
	return rep.(*pb.LoginAck), nil
}

func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Login)
	return LoginRequest{In: Login{
		Account:  req.GetAccount(),
		Password: req.GetAccount(),
	}}, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(LoginResponse)
	ack := &pb.LoginAck{}
	if resp.Err != nil {
		ack.Err = resp.Err.Error()
	} else {
		ack.Token = resp.Ack.Token
	}
	return ack, nil
}
