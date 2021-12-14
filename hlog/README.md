# 封装 zap 日志注入 trace 信息 Trace Id（内含 gin 例子）

### [hlog](https://github.com/hwholiday/learning_tools/tree/master/hlog) (源码地址)

- 实现自动切割文件 (基于 lumberjack 实现)
- 实现可传递 trace 信息 （基于 Context 实现）

###   

#### 配置

- Development bool // 是否开发模式
- LogFileDir string // 日志路径
- AppName string // APP名字
- MaxSize int //文件多大开始切分
- MaxBackups int //保留文件个数
- MaxAge int //文件保留最大实际
- Level string // 日志打印等级
- CtxKey string //通过 ctx 传递 hlog 信息
- WriteFile bool // 是否写入文件
- WriteConsole bool // 是否控制台打印

##### 实现自动切割文件核心代码

```base
 zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.opts.LogFileDir + "/" + l.opts.AppName + ".log",
		MaxSize:    l.opts.MaxSize,
		MaxBackups: l.opts.MaxBackups,
		MaxAge:     l.opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	})
```

#### 实现可传递 trace 信息核心代码

```base
func (l *Logger) GetCtx(ctx context.Context) *zap.Logger {
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
```

#### 例子（普通）

```base
hlog.NewLogger()
hlog.GetLogger().Info("hconf example success")

{"L":"INFO","T":"2021-12-14T11:43:13.276+0800","C":"hlog/zap.go:34","M":"[initLogger] zap plugin initializing completed"}
{"L":"INFO","T":"2021-12-14T11:43:13.277+0800","C":"hlog/zap_test.go:12","M":"hconf example success"}
```

#### 例子（gin）

```base
func AddTraceId() gin.HandlerFunc {
	return func(g *gin.Context) {
		traceId := g.GetHeader("traceId")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		ctx, log := hlog.GetLogger().AddCtx(g.Request.Context(), zap.Any("traceId", traceId))
		g.Request = g.Request.WithContext(ctx)
		log.Info("AddTraceId success")
		g.Next()
	}
}

log := hlog.GetLogger().GetCtx(context.Request.Context())
		log.Info("test")
		log.Debug("test")	
```

#### 例子（gin）开发模式

```base	
hlog.NewLogger()	

curl http://127.0.0.1:8888/test

{"L":"INFO","T":"2021-12-14T11:46:00.170+0800","C":"example/main.go:35","M":"hconf example success"}
{"L":"INFO","T":"2021-12-14T11:46:03.956+0800","C":"example/main.go:19","M":"AddTraceId success","traceId":"b1471a7c-5ae8-4bfd-bbdc-5312e072719c"}
{"L":"INFO","T":"2021-12-14T11:46:03.956+0800","C":"example/main.go:31","M":"test","traceId":"b1471a7c-5ae8-4bfd-bbdc-5312e072719c"}
{"L":"DEBUG","T":"2021-12-14T11:46:03.956+0800","C":"example/main.go:32","M":"test","traceId":"b1471a7c-5ae8-4bfd-bbdc-5312e072719c"}
```

#### 例子（gin）生产模式

```base
hlog.NewLogger(
	hlog.SetDevelopment(false))

curl http://127.0.0.1:8888/test
	
{"level":"info","ts":1639453661.4718382,"caller":"example/main.go:36","msg":"hconf example success"}
{"level":"info","ts":1639453664.7402327,"caller":"example/main.go:19","msg":"AddTraceId success","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
{"level":"info","ts":1639453664.7402515,"caller":"example/main.go:32","msg":"test","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
{"level":"debug","ts":1639453664.7402549,"caller":"example/main.go:33","msg":"test","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
		
```
