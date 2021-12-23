package main

import (
	"github.com/hwholiday/learning_tools/go-kit/v2/utils"
	"github.com/hwholiday/learning_tools/go-kit/v2/v2_endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v2/v2_service"
	"github.com/hwholiday/learning_tools/go-kit/v2/v2_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v2_service.NewService(utils.GetLogger())
	endpoints := v2_endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v2_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
