package v1_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"learning_tools/go-kit/v1/v1_server"
)

type ServerEndpoints struct {
	TestEndpoint endpoint.Endpoint
}

func TestEndpoint(svc v1_server.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return svc.Test(ctx, req)
	}
}

func (s ServerEndpoints) Test(cxt context.Context, req string) (res string, err error) {
	var resp interface{}
	resp, err = s.TestEndpoint(cxt, req)
	if err != nil {
		return "", err
	}
	res = resp.(string)
	return
}
