package main

import (
	"learning_tools/go-kit/v1/v1_endpoint"
	"learning_tools/go-kit/v1/v1_service"
	"learning_tools/go-kit/v1/v1_transport"
	"net/http"
)

func main() {
	server := v1_service.NewService()
	endpoints := v1_endpoint.NewEndPointServer(server)
	httpHandler := v1_transport.NewHttpHandler(endpoints)
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
