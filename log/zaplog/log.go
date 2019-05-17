package zaplog

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Zap *zap.Logger

type ToolLogger struct {
	Filename   string        // Filename是要写入日志的文件。 备份日志文件将被保留
	MaxSize    int           // 一个文件多少Ｍ大于该数字开始切分文件
	MaxAge     int           // MaxAge是根据日期保留旧日志文件的最大天数
	MaxBackups int           // MaxBackups是要保留的最大旧日志文件数
	Compress   bool          //压缩确定是否应压缩旋转的日志文件
	Level      zapcore.Level //日志等级　zap.DebugLevel zap.InfoLevel  zap.ErrorLevel ... (设置为debug日志将输入到控制台和文件中,其他模式只输出到文件)
}

func InitZapLogger(tool *ToolLogger) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//hook := lumberjack.Logger{
	//	Filename:   filepath.Join(path, tool.Filename),
	//	MaxSize:    tool.MaxSize,
	//	MaxBackups: tool.MaxBackups,
	//	MaxAge:     tool.MaxAge,
	//	Compress:   tool.Compress,
	//	LocalTime:  true,
	//}
	//w := zapcore.AddSync(&hook)
	logf, err := rotatelogs.New(
		filepath.Join(path, tool.Filename)+".%Y_%m_%d_%H:%M",
		rotatelogs.WithLinkName(filepath.Join(path, tool.Filename)),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)
	w := zapcore.AddSync(logf)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= tool.Level
	})
	var core zapcore.Core
	if tool.Level == zapcore.DebugLevel {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = timeV2Encoder
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleDebugging := zapcore.Lock(os.Stdout)
		core = zapcore.NewTee(
			// 打印在控制台
			zapcore.NewCore(consoleEncoder, consoleDebugging, priority),
			// 打印在文件中
			zapcore.NewCore(consoleEncoder, w, priority),
		)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = timeUnixNanoV2Encoder
		core = zapcore.NewTee(
			// 打印在文件中
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), w, priority),
		)
	}
	Zap = zap.New(core).WithOptions(zap.AddCaller())
	Zap.Info("DefaultLogger init success")
}
func timeV2Encoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func timeUnixNanoV2Encoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano())
}

func ZapLoggerClose() {
	if Zap != nil {
		Zap.Sync()
	}
}
