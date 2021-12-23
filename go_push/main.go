package main

import (
	"github.com/hwholiday/learning_tools/go_push/gateway"
	"github.com/hwholiday/learning_tools/go_push/logic"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go gateway.InitWsServer()
	go logic.InitHttpServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
