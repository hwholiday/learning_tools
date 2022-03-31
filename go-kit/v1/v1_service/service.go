package v1_service

import "context"

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
}

func NewService() Service {
	return &baseServer{}
}

func (s baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	return AddAck{Res: in.A + in.B}
}
