package hlog

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"testing"
)

func TestGetLogger(t *testing.T) {
	NewLogger()
	GetLogger().Info("hconf example success")
	// 可以在中间件内赋值
	ctx, hlog := GetLogger().AddCtx(context.Background(), zap.String("traceId", uuid.New().String()))
	hlog.Debug("TestGetLogger", zap.Any("t", "t"))
	FA(ctx)
	FB(ctx)

	// 可以在中间件内赋值
	ctx, hlog = GetLogger().AddCtx(context.Background(), zap.String("traceId", uuid.New().String()))
	hlog.Info("TestGetLogger", zap.Any("t", "t"))
	FA(ctx)
	FB(ctx)
}

func FA(ctx context.Context) {
	hlog := GetLogger().GetCtx(ctx)
	hlog.Info("FA", zap.Any("a", "a"))
}

func FB(ctx context.Context) {
	hlog := GetLogger().GetCtx(ctx)
	hlog.Info("FB", zap.Any("b", "b"))
	FC(ctx)
}
func FC(ctx context.Context) {
	hlog := GetLogger().GetCtx(ctx)
	hlog.Info("FC", zap.Any("c", "c"))
}
