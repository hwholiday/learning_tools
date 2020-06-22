package main

import (
	"learning_tools/websocket/gateway/ws"
)

func main() {
	ws.InitWsServer()
	select {}
}
