package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func Test(t *testing.T) {
	client, err := rpc.Dial("tcp", "192.168.2.28:9023")
	if err != nil {
		panic(err)
	}
	var (
		arg   = Arg{Arg: "coco"}
		reply = Reply{}
	)
	err = client.Call("RPC.Ping", &arg, &reply)
	if err != rpc.ErrShutdown {
		client.Close()
	}
	fmt.Println(reply)
}
