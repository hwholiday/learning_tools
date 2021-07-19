package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"learning_tools/etcd/discovery"
	"learning_tools/grpc/etcd-grpc/api"
	"log"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	r := discovery.NewDiscovery(
		discovery.SetName("hwholiday.srv.msg"),
		discovery.SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"172.12.12.165:2379"},
			DialTimeout: time.Second * 5,
		}))
	resolver.Register(r)
	// 连接服务器
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", r.Scheme(), "hwholiday.srv.msg"),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "version")),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	apiClient := api.NewApiClient(conn)
	ctx := context.WithValue(context.Background(), "version", "v1")
	req, err := apiClient.ApiTest(ctx, &api.Request{Input: "v1v1v1v1v1v1v1v1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.String())
	ctx = context.WithValue(context.Background(), "version", "v2")
	req, err = apiClient.ApiTest(ctx, &api.Request{Input: "v2v2v2v2v2v2v2v2v2"})
	if err != nil {
		fmt.Println(err)
	}
	select {}

}
