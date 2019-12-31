package main

import (
	"go.uber.org/zap"
	"learning_tools/all_packaged_library/logtool"
	"learning_tools/go-kit/v2/v2_endpoint"
	"learning_tools/go-kit/v2/v2_service"
	"learning_tools/go-kit/v2/v2_transport"
	"net/http"
)

func main() {
	lg := logtool.NewLogger(
		logtool.SetAppName("test_app"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
	server := v2_service.NewService(lg)
	server = v2_service.NewLogMiddlewareServer(lg)(server)
	endpoints := v2_endpoint.NewEndPointServer(server, lg)
	httpHandler := v2_transport.NewHttpHandler(endpoints, lg)
	lg.Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
