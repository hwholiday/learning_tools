package main

import (
	"github.com/hwholiday/libs/perf"
	"github.com/hwholiday/libs/quit"
	"fmt"
	"github.com/hwholiday/libs/logtool"
	"learning_tools/tcp/network"
	"go.uber.org/zap"
)

func main() {
	perf.StartPprof([]string{"192.168.2.28:9022"})
	logtool.InitZapLogger(&logtool.ToolLogger{Filename: "logtool.log", MaxSize: 10, MaxAge: 30, MaxBackups: 30, Compress: false, Level: zap.DebugLevel})
	go network.InitTcp()
	quit.QuitSignal(func() {
		fmt.Println("开始退出")
		fmt.Println("退出程序")
	})
}
