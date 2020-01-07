package v5_transport

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"learning_tools/go-kit/v5/v5_user/v5_service"
)

type LogErrorHandler struct {
	logger *zap.Logger
}

func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(v5_service.ContextReqUUid)), zap.Error(err))
}
