package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hwholiday/learning_tools/hlog"
	"go.uber.org/zap"
	"net/http"
)

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

// curl http://127.0.0.1:8888/test
func main() {
	hlog.NewLogger(
		hlog.SetDevelopment(false))
	g := gin.New()
	g.Use(AddTraceId())
	g.GET("/test", func(context *gin.Context) {
		log := hlog.GetLogger().GetCtx(context.Request.Context())
		log.Info("test")
		log.Debug("test")
		context.JSON(200, "success")
	})
	hlog.GetLogger().Info("hconf example success")
	http.ListenAndServe(":8888", g)
}

// curl http://127.0.0.1:8888/test
//{"L":"INFO","T":"2021-12-14T11:12:24.916+0800","C":"example/main.go:35","M":"hconf example success"}
//{"L":"INFO","T":"2021-12-14T11:12:29.179+0800","C":"example/main.go:19","M":"AddTraceId success","traceId":"e92b1a38-f9ce-4d1a-8f40-c88caf735844"}
//{"L":"INFO","T":"2021-12-14T11:12:29.179+0800","C":"example/main.go:31","M":"test","traceId":"e92b1a38-f9ce-4d1a-8f40-c88caf735844"}
//{"L":"DEBUG","T":"2021-12-14T11:12:29.179+0800","C":"example/main.go:32","M":"test","traceId":"e92b1a38-f9ce-4d1a-8f40-c88caf735844"}
