package main

import (
	"learning_tools/go-kit/v3/utils"
	"learning_tools/go-kit/v3/v3_endpoint"
	"learning_tools/go-kit/v3/v3_service"
	"learning_tools/go-kit/v3/v3_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v3_service.NewService(utils.GetLogger())
	endpoints := v3_endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v3_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
