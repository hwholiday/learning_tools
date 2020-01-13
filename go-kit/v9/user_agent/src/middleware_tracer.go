package src

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"learning_tools/go-kit/v9/user_agent/pb"
)

type tracerMiddlewareServer struct {
	next   Service
	tracer opentracing.Tracer
}

func NewTracerMiddlewareServer(tracer opentracing.Tracer) NewMiddlewareServer {
	return func(service Service) Service {
		return tracerMiddlewareServer{
			next:   service,
			tracer: tracer,
		}
	}
}

func (l tracerMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, l.tracer, "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("account", in.GetAccount(), "password", in.GetPassword())
		span.Finish()
	}()
	out, err = l.next.Login(ctxContext, in)
	return
}
