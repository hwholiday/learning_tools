package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learning_tools/istio/api"
	"log"
	"net"
)

const version = "logic_v3"

type Server struct{}

func (s *Server) ReqName(c context.Context, req *api.Req) (*api.Res, error) {
	fmt.Println(version, " ReqName ", req.Name)
	return &api.Res{Name: fmt.Sprintf("%s_%s", version, req.Name)}, nil
}

func (s *Server) ReqVersion(c context.Context, req *api.Req) (*api.Res, error) {
	data, err := rd.Get(context.Background(), "istio:test").Result()
	if err != nil {
		fmt.Println("ReqVersion get", err)
		return nil, err
	}
	return &api.Res{Name: data}, nil
}

var rd *Client

func main() {
	var err error
	rd, err = NewRedis(Config{
		PoolSize: 20,
		Addr:     []string{"redis-master-1.im.svc.cluster.local:7001", "redis-master-2.im.svc.cluster.local:7002", "redis-master-3.im.svc.cluster.local:7003"},
	})
	if err != nil {
		log.Panic(err)
		return
	}
	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Panic(err)
		return
	}
	g := grpc.NewServer()
	api.RegisterNameServer(g, &Server{})
	fmt.Println("GRPC 8099 启动成功")
	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}
}
