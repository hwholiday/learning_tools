package v2_service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}

func (s baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	//模拟耗时
	time.Sleep(time.Millisecond * 2)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v2_service Service", "TestAdd 处理请求"))
	ack:=AddAck{Res: in.A + in.B}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v2_service Service", "TestAdd 处理请求"),zap.Any("处理返回值",ack))
	return ack
}
