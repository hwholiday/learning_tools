package v1_service

import "context"

type Service interface {
	TestAdd(ctx context.Context,a int) int
}

type baseServer struct {

}

func NewService() Service {
	return &baseServer{}
}

func (s baseServer) TestAdd(ctx context.Context,a int) int {
	return a +10
}
