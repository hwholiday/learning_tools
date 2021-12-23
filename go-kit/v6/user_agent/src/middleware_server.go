package src

import (
	"context"
	"fmt"
	"github.com/hwholiday/learning_tools/go-kit/v6/user_agent/pb"
	"go.uber.org/zap"
	"time"
)

const ContextReqUUid = "req_uuid"

type NewMiddlewareServer func(Service) Service

type logMiddlewareServer struct {
	logger *zap.Logger
	next   Service
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return logMiddlewareServer{
			logger: log,
			next:   service,
		}
	}
}

func (l logMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("time", time.Since(start)), zap.Any("err", err))
	}(time.Now())
	out, err = l.next.Login(ctx, in)
	return
}
