package src

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/time/rate"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

type LoginRequest struct {
	In Login `json:"in"`
}

type LoginResponse struct {
	Ack LoginAck `json:"ack"`
	Err error    `json:"err"`
}

func NewEndPointServer(svc Service, limit *rate.Limiter, tracer opentracing.Tracer) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
		loginEndPoint = NewTracerEndpointMiddleware(tracer)(loginEndPoint)

	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Login(ctx context.Context, in Login)  (ack LoginAck, err error) {
	request := LoginRequest{In: in}
	var res interface{}
	res, err = s.LoginEndPoint(ctx, request)
	if err != nil {
		return
	}
	return res.(LoginResponse).Ack, res.(LoginResponse).Err
}

func MakeLoginEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)
		res, err := s.Login(ctx, req.In)
		return LoginResponse{
			Ack: res,
			Err: err,
		}, nil
	}
}
