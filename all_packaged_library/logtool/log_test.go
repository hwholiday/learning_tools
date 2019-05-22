package logtool

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestInitZapV2Logger(t *testing.T) {
	data := &Options{}
	data.Development = true
	initLogger(data)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		GetLogger().Debug(fmt.Sprint("debug log ", i), zap.Int("line", 47))
		GetLogger().Info(fmt.Sprint("Info log ", i), zap.Any("level", "1231231231"))
		GetLogger().Warn(fmt.Sprint("warn log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		GetLogger().Error(fmt.Sprint("err log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
}
