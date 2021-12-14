package hlog

import "path/filepath"

type Options struct {
	Development  bool
	LogFileDir   string
	AppName      string
	MaxSize      int //文件多大开始切分
	MaxBackups   int //保留文件个数
	MaxAge       int //文件保留最大实际
	Level        string
	CtxKey       string //通过 ctx 传递 hlog 对象
	WriteFile    bool
	WriteConsole bool
}

type HLogOptions func(*Options)

func newOptions(opts ...HLogOptions) *Options {
	opt := &Options{
		Development:  true,
		AppName:      "hlog-app",
		MaxSize:      100,
		MaxBackups:   60,
		MaxAge:       30,
		Level:        "debug",
		CtxKey:       "hlog-ctx",
		WriteFile:    false,
		WriteConsole: true,
	}
	opt.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
	opt.LogFileDir += "/logs/"
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func SetDevelopment(development bool) HLogOptions {
	return func(options *Options) {
		options.Development = development
	}
}

func SetLogFileDir(logFileDir string) HLogOptions {
	return func(options *Options) {
		options.LogFileDir = logFileDir
	}
}

func SetAppName(appName string) HLogOptions {
	return func(options *Options) {
		options.AppName = appName
	}
}

func SetMaxSize(maxSize int) HLogOptions {
	return func(options *Options) {
		options.MaxSize = maxSize
	}
}
func SetMaxBackups(maxBackups int) HLogOptions {
	return func(options *Options) {
		options.MaxBackups = maxBackups
	}
}
func SetMaxAge(maxAge int) HLogOptions {
	return func(options *Options) {
		options.MaxAge = maxAge
	}
}

func SetLevel(level string) HLogOptions {
	return func(options *Options) {
		options.Level = level
	}
}

func SetCtxKey(ctxKey string) HLogOptions {
	return func(options *Options) {
		options.CtxKey = ctxKey
	}
}

func SetWriteFile(writeFile bool) HLogOptions {
	return func(options *Options) {
		options.WriteFile = writeFile
	}
}

func SetWriteConsole(writeConsole bool) HLogOptions {
	return func(options *Options) {
		options.WriteConsole = writeConsole
	}
}
