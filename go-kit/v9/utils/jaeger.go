package utils

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"
	"strings"
)

func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},

		ServiceName: serviceName,
	}

	tracer, closer, err = cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}

var (
	TracingComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc"}
)

type MDReaderWriter struct {
	metadata.MD
}

func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}

func JaegerClientMiddleware(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var parentCtx opentracing.SpanContext
		//先判断ctx里面有没有 span 信息
		//没有就生成一个
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}
		cliSpan := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx), //父子关系的span关系
			TracingComponentTag,            //grcp tag
			ext.SpanKindRPCClient,          //客户端 tag
		)
		defer cliSpan.Finish()
		//从context中获取metadata。md.(type) == map[string][]string
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			////如果对metadata进行修改，那么需要用拷贝的副本进行修改。
			md = md.Copy()
		}
		//定义一个carrier，下面的Inject注入数据需要用到。carrier.(type) == map[string]string
		//carrier := opentracing.TextMapCarrier{}
		mdWriter := MDReaderWriter{md}
		////将span的context信息注入到carrier中
		err := tracer.Inject(cliSpan.Context(), opentracing.TextMap, mdWriter)
		if err != nil {
			grpclog.Errorf("inject to metadata err %v", err)
		}
		////创建一个新的context，把metadata附带上
		ctx = metadata.NewOutgoingContext(ctx, md)
		err = invoker(ctx, method, req, resp, cc, opts...)
		if err != nil {
			cliSpan.LogFields(log.String("err", err.Error()))
		}
		return err
	}
}

func JaegerServerMiddleware(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err %v", err)
		}
		serverSpan := tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(spanContext),
			TracingComponentTag,
			ext.SpanKindRPCServer,
		)
		defer serverSpan.Finish()
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		return handler(ctx, req)
	}
}
