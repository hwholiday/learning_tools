package main

import (
	"testing"
	"google.golang.org/grpc"
	"log"
	pr "learning_tools/grpc/server_side_streaming_rpc/proto"
	"golang.org/x/net/context"
	"io"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	client := pr.NewHowieServerClient(conn)
	server, err := client.Howie(context.Background(), &pr.Request{Input: "go"})
	if err != nil {
		log.Panic(err)
	}
	for {
		data, err := server.Recv()
		if err == io.EOF {
			log.Println("服务端停止发送")
			break
		}
		if err != nil {
			log.Print(err.Error())
		}
		log.Println(data.Output)
	}
}
