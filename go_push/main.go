package main

import (
	"learning_tools/go_push/gateway"
	"learning_tools/go_push/logic"
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
