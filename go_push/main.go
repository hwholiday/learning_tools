package main

import "learning_tools/go_push/gateway"

func main() {
	gateway.InitWsServer()
	gateway.InitHttpServer()
}
