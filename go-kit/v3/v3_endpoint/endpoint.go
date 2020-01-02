package v3_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"learning_tools/go-kit/v3/v3_service"
)

type EndPointServer struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v3_service.Service, log *zap.Logger) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
	}
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint, LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Add(ctx context.Context, in v3_service.Add) v3_service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(v3_service.AddAck)
}

func MakeAddEndPoint(s v3_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v3_service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

func (s EndPointServer) Login(ctx context.Context, in v3_service.Login) (v3_service.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	return res.(v3_service.LoginAck), err
}

func MakeLoginEndPoint(s v3_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v3_service.Login)
		return s.Login(ctx, req)
	}
}
