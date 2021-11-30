package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/adpter/http/auth_handles"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/adpter/http/routers"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/service"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/conf"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewHttp(s *conf.AppConfig, auth service.AuthSrv) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	h := auth_handles.NewHandles(s, auth)
	routers.SetRouters(g, h)
	server := &http.Server{
		Addr:           s.NetConf.ServerAddr,
		Handler:        g,
		ReadTimeout:    time.Duration(s.NetConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(s.NetConf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.GetLogger().Info("auth server start success", zap.Any("addr", s.NetConf.ServerAddr))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
