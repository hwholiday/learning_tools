package utils

import (
	"go.uber.org/zap"
	"learning_tools/all_packaged_library/logtool"
)

var logger *zap.Logger

func NewLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("test_app"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
