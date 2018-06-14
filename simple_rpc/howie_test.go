package main

import (
	"testing"
	"google.golang.org/grpc"
	"log"
	"test/simple_rpc/proto"
	"golang.org/x/net/context"
	"fmt"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8099",grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	c:=proto.NewHowieClient(conn)
	info,err:=c.LoL(context.Background(),&proto.HowieUp{Name:"howie"})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(info.Msg)
}
