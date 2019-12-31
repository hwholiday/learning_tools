package v2_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"learning_tools/go-kit/v2/v2_service"
)

type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v2_service.Service, log *zap.Logger) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}

func (s EndPointServer) Add(ctx context.Context, in v2_service.Add) v2_service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(v2_service.AddAck)
}

func MakeAddEndPoint(s v2_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v2_service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}
