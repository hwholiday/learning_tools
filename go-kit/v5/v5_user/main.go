package main

import (
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"learning_tools/go-kit/v5/utils"
	"learning_tools/go-kit/v5/v5_user/v5_endpoint"
	"learning_tools/go-kit/v5/v5_user/v5_service"
	"learning_tools/go-kit/v5/v5_user/v5_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1) //每秒产生10个令牌,令牌桶的可以装1个令牌
	uberLimit := ratelimit.New(1)         //一秒请求一次
	server := v5_service.NewService(utils.GetLogger())
	endpoints := v5_endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	httpHandler := v5_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
