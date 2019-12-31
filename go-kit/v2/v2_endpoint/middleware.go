package v2_endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"learning_tools/go-kit/v2/v2_service"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprintln(ctx.Value(v2_service.ContextReqUUid)), zap.Any("time", time.Since(begin).String()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
