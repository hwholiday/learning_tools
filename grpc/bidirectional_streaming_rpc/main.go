package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	pr "learning_tools/grpc/bidirectional_streaming_rpc/proto"
	"io"
	"time"
)

type Server struct{}

func (s *Server) Chat(stream pr.ChatService_ChatServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			request, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("客户端发送结束:", err.Error())
				return nil
			}
			if err != nil {
				fmt.Println("接受数据出错:", err.Error())
				return err
			}
			switch request.Input {
			case "结束":
				fmt.Println("结束该流程")
				if err := stream.Send(&pr.Response{Output: "结束聊天"}); err != nil {
					return nil
				}
				return nil

			case "聊天":
				fmt.Println("客户端想要聊天")
				if err := stream.Send(&pr.Response{Output: "来聊天啊:" + time.Now().Format("2006-01-02 15:04:05")}); err != nil {
					return nil
				}

			default:
				if err := stream.Send(&pr.Response{Output: "默认消息"}); err != nil {
					return nil
				}
			}
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Println(err)
	}
	g := grpc.NewServer()
	pr.RegisterChatServiceServer(g, &Server{})
	fmt.Println("服务启动成功")
	g.Serve(listen)

}
