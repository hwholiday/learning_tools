package main

import (
	"github.com/hwholiday/learning_tools/go-kit/v4/utils"
	"github.com/hwholiday/learning_tools/go-kit/v4/v4_endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v4/v4_service"
	"github.com/hwholiday/learning_tools/go-kit/v4/v4_transport"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
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
