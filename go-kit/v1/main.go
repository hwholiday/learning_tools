package main

import (
	"fmt"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_service"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_transport"
	"net/http"
)

func main() {
	server := v1_service.NewService()
	endpoints := v1_endpoint.NewEndPointServer(server)
	httpHandler := v1_transport.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
