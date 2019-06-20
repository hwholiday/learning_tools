package zaplog

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestInitZapV2Logger(t *testing.T) {
	InitZapLogger(&ToolLogger{Filename: "logtool.log", MaxSize: 3, MaxAge: 30, MaxBackups: 30, Compress: false, Level: zap.InfoLevel})
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		Zap.Info(fmt.Sprint("test log ", i), zap.Int("line", 47))
		Zap.Debug(fmt.Sprint("debug log ", i), zap.Any("level", "1231231231"))
		Zap.Info(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		Zap.Warn(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
}

func TestZapLog(t *testing.T) {
	data:=&Options{}
	data.Development=true
	initLogger(data)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		GetLogger().Debug(fmt.Sprint("debug log ", i), zap.Int("line", 47))
		GetLogger().Info(fmt.Sprint("Info log ", i), zap.Any("level", "1231231231"))
		GetLogger().Warn(fmt.Sprint("warn log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		GetLogger().Error(fmt.Sprint("err log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
}
