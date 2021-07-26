package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"learning_tools/etcd/discovery"
	"learning_tools/etcd/etcd-grpc/api"
	"log"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	r, err := discovery.NewDiscovery(
		discovery.SetName("hwholiday.srv.app"),
		discovery.SetLoadBalancingPolicy(discovery.VersionLB),
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
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, discovery.VersionLB)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	apiClient := api.NewApiClient(conn)
	ctx := context.WithValue(context.Background(), "version", "v1")
	res, err := apiClient.ApiTest(ctx, &api.Request{Input: "v1v1v1v1v1"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Output)
}
