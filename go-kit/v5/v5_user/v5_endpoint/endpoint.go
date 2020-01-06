package v5_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v5/v5_user/v5_service"
)

type EndPointServer struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v5_service.Service, log *zap.Logger, limit *rate.Limiter, limiter ratelimit.Limiter) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
		addEndPoint = NewUberRateMiddleware(limiter)(addEndPoint)
	}
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint, LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Add(ctx context.Context, in v5_service.Add) v5_service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(v5_service.AddAck)
}

func MakeAddEndPoint(s v5_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v5_service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

func (s EndPointServer) Login(ctx context.Context, in v5_service.Login) (v5_service.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	return res.(v5_service.LoginAck), err
}

func MakeLoginEndPoint(s v5_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v5_service.Login)
		return s.Login(ctx, req)
	}
}
