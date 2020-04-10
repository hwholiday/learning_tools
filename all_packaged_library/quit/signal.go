package quit

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

//syscall.SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
//syscall.SIGTERM 结束程序(可以被捕获、阻塞或忽略)
//syscall.SIGINT 用户发送INTR字符(Ctrl+C)触发
func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("server start success pid:%d\n", os.Getpid())
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}
