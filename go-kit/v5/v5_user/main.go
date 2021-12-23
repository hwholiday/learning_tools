package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/hwholiday/learning_tools/go-kit/v5/utils"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/pb"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/v5_endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/v5_service"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/v5_transport"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)
	server := v5_service.NewService(utils.GetLogger())
	endpoints := v5_endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit)
	grpcServer := v5_transport.NewGRPCServer(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run :8881")
	grpcListener, err := net.Listen("tcp", ":8881")
	if err != nil {
		utils.GetLogger().Warn("Listen", zap.Error(err))
		os.Exit(0)
	}
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		utils.GetLogger().Warn("Serve", zap.Error(err))
		os.Exit(0)
	}

}
