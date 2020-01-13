package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/time/rate"
)

func NewGolangRateAllowMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return "", errors.New("limit req  Allow")
			}
			return next(ctx, request)
		}
	}
}

func NewTracerEndpointMiddleware(tracer opentracing.Tracer) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "endpoint", opentracing.Tag{
				Key:   string(ext.Component),
				Value: "NewTracerEndpointMiddleware",
			})
			defer span.Finish()
			return next(ctxContext, request)
		}
	}
}
