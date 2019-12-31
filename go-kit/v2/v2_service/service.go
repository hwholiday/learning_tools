package v2_service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	return &baseServer{logger: log}
}

func (s baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("v2_service Service", "TestAdd"))
	return AddAck{Res: in.A + in.B}
}
