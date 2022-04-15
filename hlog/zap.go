package hlog

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l            *Logger
	outWrite     zapcore.WriteSyncer       // IO输出
	debugConsole = zapcore.Lock(os.Stdout) // 控制台标准输出
	once         sync.Once
)

type Logger struct {
	*zap.Logger
	opts      *Options
	zapConfig zap.Config
}

func NewLogger(opts ...HLogOptions) {
	once.Do(func() {
		l = &Logger{
			opts: newOptions(opts...),
		}
		l.loadCfg()
		l.initHLog()
		l.Info("[initLogger] zap plugin initializing completed")
	})
}

// GetLogger returns logger
func GetLogger() *Logger {
	if l == nil {
		fmt.Println("Please initialize the hlog service first")
		return nil
	}
	return l
}

func (l *Logger) GetCtx(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(l.opts.CtxKey).(*zap.Logger)
	if ok {
		return log
	}
	return l.Logger
}

func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(l.opts.CtxKey).(*zap.Logger)
	if ok {
		return log
	}
	return l.Logger
}

func (l *Logger) AddCtx(ctx context.Context, field ...zap.Field) (context.Context, *zap.Logger) {
	log := l.With(field...)
	ctx = context.WithValue(ctx, l.opts.CtxKey, log)
	return ctx, log
}

func (l *Logger) initHLog() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}
func (l *Logger) GetLevel() (level zapcore.Level) {
	switch strings.ToLower(l.opts.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel //默认为调试模式
	}
}

func (l *Logger) loadCfg() {
	if l.opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
		//l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		//l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}
}

func (l *Logger) setSyncers() {
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.opts.LogFileDir + "/" + l.opts.AppName + ".log",
		MaxSize:    l.opts.MaxSize,
		MaxBackups: l.opts.MaxBackups,
		MaxAge:     l.opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	})
	return
}

func (l *Logger) cores() zap.Option {
	encoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l.GetLevel()
	})
	var cores []zapcore.Core
	if l.opts.WriteFile {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(encoder, outWrite, priority),
		}...)
	}
	if l.opts.WriteConsole {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(encoder, debugConsole, priority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

// 可自定义时间
//func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
//}
//
//func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendInt64(t.UnixNano() / 1e6)
//}
