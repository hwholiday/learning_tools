package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	initLogger(&Options{
		LogFileDir: "logs",
		AppName:    "gid",
		Level:      zapcore.InfoLevel,
	})
	log := GetLogger()
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second / 2)
		log.Info(fmt.Sprint("test log ", i), zap.Int("line", 47))
		log.Debug(fmt.Sprint("debug log ", i), zap.Any("level", "1231231231"))
		log.Error(fmt.Sprint("error log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		log.Warn(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
}
