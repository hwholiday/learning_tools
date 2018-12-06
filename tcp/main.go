package main

import (
	"github.com/hwholiday/libs/perf"
	"github.com/hwholiday/libs/quit"
	"fmt"
	"github.com/hwholiday/libs/logtool"
	"learning_tools/tcp/network"
)

func main() {
	perf.StartPprof([]string{"192.168.2.28:9022"})
	logtool.InitZapLogger("ghost.log", true)
	go network.InitTcp()
	quit.QuitSignal(func() {
		fmt.Println("开始退出")
		fmt.Println("退出程序")
	})
}
