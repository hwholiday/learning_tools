package v1_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"learning_tools/go-kit/v1/v1_service"
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

func (s EndPointServer) Add(ctx context.Context, a int) int {
	res, _ := s.AddEndPoint(ctx, a)
	return res.(int)
}

func MakeAddEndPoint(s v1_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(int)
		res := s.TestAdd(ctx, req)
		//return res, errors.New("123123")
		return res, nil
	}
}
