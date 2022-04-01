package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hwholiday/learning_tools/hlb-grpc/discovery"
	"github.com/hwholiday/learning_tools/hlb-grpc/example/api"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func TestClient(t *testing.T) {
	r, err := discovery.NewDiscovery(
		discovery.SetName("hwholiday.srv.app"),
		discovery.SetLoadBalancingPolicy(discovery.CustomizeLB),
		//discovery.SetVersion("v1"),
		discovery.SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"172.12.12.165:2379"},
			DialTimeout: time.Second * 5,
		}))
	if err != nil {
		panic(err)
	}
	resolver.Register(r)
	// 连接服务器
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", r.Scheme(), ""),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, discovery.CustomizeLB)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	apiClient := api.NewApiClient(conn)
	ctx := discovery.BuildCtxFilter(context.Background(), map[string]string{
		"gateway": "1",
	})
	for i := 0; i < 10000000; i++ {
		time.Sleep(time.Second / 10)
		res, err := apiClient.ApiTest(ctx, &api.Request{Input: "v1v1v1v1v1"})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.Output)
		}
	}
}
