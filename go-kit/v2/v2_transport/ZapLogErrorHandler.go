package v2_transport

import (
	"context"
	"fmt"
	"github.com/hwholiday/learning_tools/go-kit/v2/v2_service"
	"go.uber.org/zap"
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
	h.logger.Warn(fmt.Sprint(ctx.Value(v2_service.ContextReqUUid), zap.Error(err)))
}
