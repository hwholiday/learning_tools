package v5_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v5/v5_user/pb"
	"learning_tools/go-kit/v5/v5_user/v5_service"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v5_service.Service, log *zap.Logger, limit *rate.Limiter) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err!=nil{
		return nil,err
	}
	return res.(*pb.LoginAck), nil
}

func MakeLoginEndPoint(s v5_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
