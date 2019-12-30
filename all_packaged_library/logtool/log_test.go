package logtool

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestInitZapV2Logger(t *testing.T) {
	lg := NewLogger(SetAppName("test_app"), SetDevelopment(true), SetLevel(zap.DebugLevel), SetErrorFileName("error_e_e_e_e.log"))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		lg.Debug(fmt.Sprint("debug log ", 1), zap.Int("line", 47))
		lg.Info(fmt.Sprint("Info log ", 2), zap.Any("level", "1231231231"))
		lg.Warn(fmt.Sprint("warn log ", 3), zap.String("level", `{"a":"4","b":"5"}`))
		lg.Error(fmt.Sprint("err log ", 4), zap.String("level", `{"a":"7","b":"8"}`))

	}
}
