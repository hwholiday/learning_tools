package utils

import (
	"github.com/hwholiday/learning_tools/all_packaged_library/logtool"
	"go.uber.org/zap"
)

const ContextReqUUid = "req_uuid"

var logger *zap.Logger

func NewLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("go-kit-v11-server"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
