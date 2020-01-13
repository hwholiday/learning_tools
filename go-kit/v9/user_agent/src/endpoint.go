package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v9/user_agent/pb"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc Service, limit *rate.Limiter,tracer opentracing.Tracer) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
		loginEndPoint = NewTracerEndpointMiddleware(tracer)(loginEndPoint)

	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		fmt.Println("s.LoginEndPoint", err)
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

func MakeLoginEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
