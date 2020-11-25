package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
	"testing"
	"time"
)

func InitJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHostPort,
		},
		ServiceName: serviceName,
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func node2(data map[string]string) {
	tracer, closer := InitJaegerTracer("jaeger-test", "172.13.3.160:6831")
	defer closer.Close()
	spanContext, err := tracer.Extract(opentracing.TextMap, spanTextMap(data))
	if err != nil {
		panic(err)
	}
	// 创建server span
	/*serverSpan := tracer.StartSpan(
		"span_child",
		ext.RPCServerOption(spanContext),
		ext.SpanKindRPCServer,
	)*/
	serverSpan := tracer.StartSpan(
		"span_child",
		opentracing.ChildOf(spanContext),
	)
	serverSpan.SetTag("node", "node2")
	defer serverSpan.Finish()
	return
}

func TestJaegerTest(t *testing.T) {
	tracer, closer := InitJaegerTracer("jaeger-test", "172.13.3.160:6831")
	defer closer.Close()
	span := tracer.StartSpan("span_root")
	span.SetTag("node", "node1")
	spanData := make(map[string]string)
	carrier := spanTextMap(spanData)
	if err := span.Tracer().Inject(span.Context(), opentracing.TextMap, carrier); err != nil {
		panic(err)
	}
	go node2(carrier)
	span.Finish()
	time.Sleep(time.Minute)
}

type spanTextMap map[string]string

func (c spanTextMap) ForeachKey(handler func(key, val string) error) error {
	for k, v := range c {
		if err := handler(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (c spanTextMap) Set(key, val string) {
	c[key] = val
}
