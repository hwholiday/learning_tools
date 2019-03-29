package main

import (
	"log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"errors"
)

var Logger *zap.Logger

func InitLogger() {
	isDebug := true
	initLogger(isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func initLogger(isDebug bool) {
	var cfg zap.Config
	if isDebug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Development = true
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg.Development = false
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"stdout", "./out.log"}
		cfg.ErrorOutputPaths = []string{"stderr"}
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	}
	cfg.EncoderConfig.EncodeTime = TimeEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
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
