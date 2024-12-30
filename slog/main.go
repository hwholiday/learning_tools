package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

// 定义一个新的类型作为 context 的键，以避免键冲突
type contextKey string

const loggerKey = contextKey("logger")

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				return slog.Attr{
					Key:   a.Key,
					Value: slog.Int64Value(time.Now().Unix()),
				}
			}
			return a
		},
	})).With(slog.String("log_app_name", "test")))
	logger := slog.With(slog.String("path", "/ping"))
	logger.Error("failed", slog.Int("aaa", 2))
	ctx := context.WithValue(context.Background(), loggerKey, logger)
	slog.Debug("aaaa")
	testCtx(ctx)
}

func testCtx(ctx context.Context) {
	logger := getLoggerByCtx(ctx)
	slog.ErrorContext(ctx, "testCtx")
	logger.Debug("testCtx")
}

func getLoggerByCtx(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}
