package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	test_agent "learning_tools/micro_v2"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	micReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(micro.Name("srv.test.client"), micro.Registry(micReg))
	service.Init()
	agent := test_agent.NewTestService("srv.test", service.Client())

	var opss client.CallOption = func(o *client.CallOptions) {
		o.RequestTimeout = time.Second * 30
		o.DialTimeout = time.Second * 30
	}

	info, err := agent.RpcUserInfo(context.TODO(), &test_agent.ReqMsg{
		UserName: "test user",
	}, opss)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(info)
}
