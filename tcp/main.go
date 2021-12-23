package main

import (
	"fmt"

	"github.com/hwholiday/learning_tools/all_packaged_library/perf"
	"github.com/hwholiday/learning_tools/all_packaged_library/quit"
	"github.com/hwholiday/learning_tools/tcp/network"
)

func main() {
	perf.StartPprof([]string{"192.168.2.28:9022"})
	go network.InitTcp()
	quit.QuitSignal(func() {
		fmt.Println("开始退出")
		fmt.Println("退出程序")
	})
}
