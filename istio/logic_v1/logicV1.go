package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learning_tools/istio/api"
	"log"
	"net"
)

const version = "logic_v1"

type Server struct{}

func (s *Server) ReqName(c context.Context, req *api.Req) (*api.Res, error) {
	fmt.Println(version, " ReqName ", req.Name)
	return &api.Res{Name: fmt.Sprintf("%s_%s", version, req.Name)}, nil
}

func (s *Server) ReqVersion(c context.Context, req *api.Req) (*api.Res, error) {
	return &api.Res{Name: version}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Panic(err)
	}
	g := grpc.NewServer()
	api.RegisterNameServer(g, &Server{})
	fmt.Println("GRPC 8099 启动成功")
	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}
}
