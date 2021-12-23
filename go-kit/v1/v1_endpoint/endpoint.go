package v1_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_service"
)

type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v1_service.Service) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}

func (s EndPointServer) Add(ctx context.Context, in v1_service.Add) v1_service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(v1_service.AddAck)
}

func MakeAddEndPoint(s v1_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v1_service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}
