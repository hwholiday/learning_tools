package main

import (
	"testing"
	"google.golang.org/grpc"
	"log"
	"test/bidirectional_streaming_rpc/proto"
	"golang.org/x/net/context"
	"io"
	"fmt"
	"bufio"
	"os"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	proto.RegisterChatServiceServer()
	client := proto.NewChatServiceClient(conn)
	ctx := context.Background()
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Panic(err)
	}

	go func() {
		/*fmt.Println("请输入消息......")
		reader := bufio.NewReader(os.Stdin)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				return
			}*/
			if err := stream.Send(&proto.Request{Input: "聊天"}); err != nil {
				return
			}
		/*}*/
	}()

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("服务端结束通知:", err.Error())
			break
		}
		if err != nil {
			fmt.Println("接受数据错误:", err.Error())
		}
		fmt.Println("服务端返回:", response.Output)
	}
}

func TestInput(t *testing.T)  {
	reader:=bufio.NewReader(os.Stdin)
	for{
		data,err:=reader.ReadString('\n')
		if err!=nil{
			continue
		}
		fmt.Println(data)
	}
}