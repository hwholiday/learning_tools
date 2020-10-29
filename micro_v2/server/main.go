package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	test_agent "micro_v2"
	"micro_v2/ecode"
	"os"
	"time"
)

type Agent struct {
}

func (a *Agent) RpcUserInfo(ctx context.Context, in *test_agent.ReqMsg, out *test_agent.ResMsg) error {
	out.Error = &test_agent.Error{
		Code: 200,
	}
	out.Info = "test success"
	fmt.Println("server  RpcUserInfo", "in", in, "out", out)
	return ecode.ParamErr
}

func main() {
	micReg := etcd.NewRegistry(registry.Addrs("172.13.3.160:2379"))
	md := server.DefaultOptions().Metadata
	md["gateway"] = "a"
	service := grpc.NewService(
		service.Name("srv.test"),
		service.Registry(micReg),
		service.Address("127.0.0.1:8088"),
		service.RegisterTTL(time.Second*10),
		service.RegisterInterval(time.Second*10),
		service.Metadata(md),
		service.WrapHandler(ecode.ServerEcodeWrapper("srv.test")),
	)
	service.Init()
	if err := test_agent.RegisterTestHandler(service.Server(), &Agent{}); err != nil {
		fmt.Println("1", err.Error())
		os.Exit(0)
	}
	if err := service.Run(); err != nil {
		fmt.Println("2", err.Error())
		os.Exit(0)
	}
}
