package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/labstack/echo/v4"
)

func main() {
	eng := NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.serveHTTP,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

// HTTP地址
func (eng *Engine) serveHTTP() error {
	server := xecho.StdConfig("http").Build()
	server.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, "Gopher Wuhan")
	})
	return eng.Serve(server)
}
