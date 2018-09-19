package main

import (
	"encoding/json"
	"fmt"
	"log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"github.com/kataras/go-errors"
)

var Logger *zap.Logger

func InitLogger() {
	lp := "out.log"
	lv := "DEBUG"
	isDebug := false
	initLogger(lp, lv, isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func initLogger(lp string, lv string, isDebug bool) {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "console",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, lv)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}
	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = TimeEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
func main() {
	InitLogger()
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	Logger.Info("info:", zap.Error(errors.New("123123")))
	Logger.Error("info:", zap.String("s", s[1]))
	Logger.Debug("info:", zap.String("s", s[2]))
	Logger.Fatal("info:", zap.String("s", s[3]))
}
