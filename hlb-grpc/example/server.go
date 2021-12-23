package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hwholiday/learning_tools/hlb-grpc/example/api"
	"github.com/hwholiday/learning_tools/hlb-grpc/register"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type ApiService struct{}

func (a ApiService) ApiTest(ctx context.Context, request *api.Request) (*api.Response, error) {
	fmt.Println(request.String())
	return &api.Response{Output: "v1v1v1v1v1v1v1v1v1v1"}, nil
}

var Addr = "0.0.0.0:8089"

func main() {
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	var srv = &ApiService{}
	api.RegisterApiServer(grpcServer, srv)
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
	s, err := register.NewRegister(
		register.SetName("hwholiday.srv.app"),
		register.SetAddress(Addr),
		//register.SetVersion("v1"),
		register.SetSrv(srv),
		register.SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"172.12.12.165:2379"},
			DialTimeout: time.Second * 5,
		}),
		register.SetMetadata(map[string]string{
			"gateway": "1",
		}),
	)
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	go func() {
		if s.ListenKeepAliveChan() {
			c <- syscall.SIGQUIT
		}
	}()
	fmt.Println("启动成功 === > ", Addr)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for a := range c {
		switch a {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出")
			fmt.Println(s.Close())
			return
		default:
			return
		}
	}

}
