package main

import (
	"github.com/hwholiday/learning_tools/websocket/gateway/ws"
)

func main() {
	ws.InitWsServer()
	select {}
}
