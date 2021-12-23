package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"github.com/hwholiday/learning_tools/go-kit/v11/utils"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Service interface {
	Login(ctx context.Context, in Login) (ack LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginAck struct {
	Token string `json:"token"`
}

func NewService(log *zap.Logger, counter metrics.Counter, histogram metrics.Histogram, tracer opentracing.Tracer) Service {
	var server Service
	server = &baseServer{log}
	server = NewTracerMiddlewareServer(tracer)(server)
	server = NewMetricsMiddlewareServer(counter, histogram)(server)
	server = NewLogMiddlewareServer(log)(server)
	return server
}

func (s baseServer) Login(ctx context.Context, in Login) (ack LoginAck, err error) {
	if in.Account != "hwholiday" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	//模拟耗时
	//rand.Seed(time.Now().UnixNano())
	//sl := rand.Int31n(10-1) + 1
	//time.Sleep(time.Duration(sl) * time.Millisecond * 100)
	//模拟错误
	/*if rand.Intn(10) > 3 {
		err = errors.New("服务器运行错误")
		return
	}*/
	ack = LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	return
}
