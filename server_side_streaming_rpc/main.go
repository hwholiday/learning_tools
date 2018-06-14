package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pr "test/server_side_streaming_rpc/proto"
	"time"
)

type Server struct{}

func (s *Server) Howie(q *pr.Request, stream pr.HowieServer_HowieServer) error {
	var i = 1
	for {
		i++
		time.Sleep(time.Second)
		if err := stream.Send(&pr.Response{Output: time.Now().Format("2006-01-02 15:04:05") + ":" + q.Input}); err != nil {
			return err
		}
		if i == 10 {
			return nil
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	g := grpc.NewServer()
	pr.RegisterHowieServerServer(g, &Server{})
	g.Serve(listen)
}
