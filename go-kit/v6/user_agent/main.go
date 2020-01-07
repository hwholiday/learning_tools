package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"learning_tools/go-kit/v6/user_agent/pb"
	"learning_tools/go-kit/v6/user_agent/src"
	"learning_tools/go-kit/v6/utils"
	"net"
	"os"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)
	server := src.NewService(utils.GetLogger())
	endpoints := src.NewEndPointServer(server, golangLimit)
	grpcServer := src.NewGRPCServer(endpoints, utils.GetLogger())
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
