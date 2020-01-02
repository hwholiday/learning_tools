package main

import (
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v4/utils"
	"learning_tools/go-kit/v4/v4_endpoint"
	"learning_tools/go-kit/v4/v4_service"
	"learning_tools/go-kit/v4/v4_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1) //每秒产生10个令牌,令牌桶的可以装1个令牌
	uberLimit := ratelimit.New(1)         //一秒请求一次
	server := v4_service.NewService(utils.GetLogger())
	endpoints := v4_endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	httpHandler := v4_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
