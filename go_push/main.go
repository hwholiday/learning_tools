package main

import (
	"learning_tools/go_push/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go gateway.InitWsServer()
	go gateway.InitHttpServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
