package main

import (
	"github.com/hwholiday/libs/perf"
	"github.com/hwholiday/libs/quit"
	"fmt"
	"net/rpc"
	"net"
)

// RPC
type RPC struct {
	auther string
}

type Arg struct {
	Arg string
}

type Reply struct {
	Reply string
}

func main() {
	perf.StartPprof([]string{"192.168.2.28:9022"})
	rpc.Register(&RPC{auther: "111111"})
	go rpcListen("tcp", "192.168.2.28:9023")
	quit.QuitSignal(func() {
		fmt.Println("开始退出")
		fmt.Println("退出程序")
	})
}

func rpcListen(network, addr string) {
	l, err := net.Listen(network, addr)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := l.Close(); err != nil {
			panic(err)
		}
	}()
	rpc.Accept(l)
}

func (r *RPC) Ping(arg *Arg, reply *Reply) error {
	reply.Reply=arg.Arg+"howie"
	return nil
}
