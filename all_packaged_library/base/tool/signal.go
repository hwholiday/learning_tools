package tool

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("server start success pid:%d\n", os.Getpid())
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}
