package main

import (
	"context"
	"net"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stderr).WithAttrs([]slog.Attr{slog.String("name", "slog-test")})))
	slog.Debug("test debug", "1")
	slog.Debug("test debug", "2")
	slog.Default().Debug("test debug", "3")
	slog.Debug("test debug", "4")
	slog.Info("test info", slog.Any("order", map[string]interface{}{"t": "2"}), slog.Int("int", 1))
	slog.Warn("test warn", slog.Any("order", map[string]interface{}{"t": "2"}), slog.Group("memory",
		slog.Int("current", 50),
		slog.Int("min", 20),
		slog.Int("max", 80)), slog.Int("int", 1))
	slog.Error("test error", net.ErrClosed, slog.Any("order", map[string]interface{}{"t": "2"}), slog.Int("int", 1))
	slog.LogAttrs(slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed), slog.Any("order", map[string]interface{}{"t": "2"}))
	logger := slog.Default().With(slog.String("key", "ididididididididididdidid"))
	logger.Debug("11111")
	logger.Context()
	ctx := context.WithValue(context.Background(), "vvvv", 1110)
	testCtx(ctx)
}

func testCtx(ctx context.Context) {
	logger := slog.Default().WithContext(ctx)
	logger.Debug("testCtx", slog.Any("key", logger.Context().Value("vvvv")))
}
