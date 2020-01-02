package v4_endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v4/utils"
	"learning_tools/go-kit/v4/v4_service"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v4_service.ContextReqUUid)), zap.Any("调用 v4_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
func AuthMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(utils.JWT_CONTEXT_KEY))
			if token == "" {
				err = errors.New("请登录")
				logger.Debug(fmt.Sprint(ctx.Value(v4_service.ContextReqUUid)), zap.Any("[AuthMiddleware]", "token == empty"), zap.Error(err))
				return "", err
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				logger.Debug(fmt.Sprint(ctx.Value(v4_service.ContextReqUUid)), zap.Any("[AuthMiddleware]", "ParseToken"), zap.Error(err))
				return "", err
			}
			if v, ok := jwtInfo["Name"]; ok {
				ctx = context.WithValue(ctx, "name", v)
			}
			return next(ctx, request)
		}
	}

}

func NewGolangRateWaitMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err = limit.Wait(ctx); err != nil {
				return "", errors.New("limit req  Wait")
			}
			return next(ctx, request)
		}
	}
}

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

func NewUberRateMiddleware(limit ratelimit.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take()
			return next(ctx, request)
		}
	}
}
