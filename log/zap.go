package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Options struct {
	LogFileDir string
	AppName    string
	MaxSize    int //文件多大开始切分
	MaxBackups int //保留文件个数
	MaxAge     int //文件保留最大实际
	Level      zapcore.Level
	zap.Config
}

var (
	l              *Logger
	sp             = string(filepath.Separator)
	outWrite       zapcore.WriteSyncer       // IO输出
	debugConsoleWS = zapcore.Lock(os.Stdout) // 控制台标准输出
)

func init() {
	l = &Logger{
		Opts: &Options{},
	}
}

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts      *Options `json:"opts"`
	zapConfig zap.Config
	inited    bool
}

func initLogger(cf ...*Options) {
	l.Lock()
	defer l.Unlock()
	if l.inited {
		l.Info("[initLogger] logger Inited")
		return
	}
	if len(cf) > 0 {
		l.Opts = cf[0]
	}
	l.loadCfg()
	l.init()
	l.Info("[initLogger] zap plugin initializing completed")
	l.inited = true
}

// GetLogger returns logger
func GetLogger() (ret *Logger) {
	return l
}

func (l *Logger) init() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) loadCfg() {
	if l.Opts.Level == zapcore.DebugLevel {
		l.zapConfig = zap.NewDevelopmentConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}
	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}
	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}
	// 默认输出到程序运行目录的logs子目录
	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}
	if l.Opts.AppName == "" {
		l.Opts.AppName = "app"
	}
	if l.Opts.MaxSize == 0 {
		l.Opts.MaxSize = 100
	}
	if l.Opts.MaxBackups == 0 {
		l.Opts.MaxBackups = 60
	}
	if l.Opts.MaxAge == 0 {
		l.Opts.MaxAge = 30
	}
}

func (l *Logger) setSyncers() {
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + ".log",
		MaxSize:    l.Opts.MaxSize,
		MaxBackups: l.Opts.MaxBackups,
		MaxAge:     l.Opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	})
	return
}

func (l *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l.Opts.Level
	})
	var cores []zapcore.Core
	if l.Opts.Level == zapcore.DebugLevel {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, debugConsoleWS, priority),
		}...)
	} else {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(fileEncoder, outWrite, priority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
