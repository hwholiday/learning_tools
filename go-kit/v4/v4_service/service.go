package v4_service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"learning_tools/go-kit/v4/utils"
	"time"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
	Login(ctx context.Context, in Login) (ack LoginAck, err error)
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
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v4_service Service", "TestAdd 处理请求"), zap.Any("请求用户", fmt.Sprint(ctx.Value("name"))))
	ack := AddAck{Res: in.A + in.B}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v4_service Service", "TestAdd 处理请求"), zap.Any("处理返回值", ack))
	return ack
}

func (s baseServer) Login(ctx context.Context, in Login) (ack LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v4_service Service", "Login 处理请求"))
	if in.Account != "hwholiday" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v4_service Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}
