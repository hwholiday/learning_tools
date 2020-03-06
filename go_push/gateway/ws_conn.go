package gateway

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WsConnection struct {
	mu        sync.Mutex
	connId    uint64
	ws        *websocket.Conn
	readChan  chan *WSMessage
	writeChan chan *WSMessage
	closeChan chan bool
	isClosed  bool
	clientIp  string
}

type WSMessage struct {
	Type int
	Data []byte
}

func (w *WsConnection) read() {
	var (
		Type int
		Data []byte
		err  error
	)

	for {
		if Type, Data, err = w.ws.ReadMessage(); err != nil {
			w.c
		}
	}
}

func (w *WsConnection)close() {
	_ = w.ws.Close()

}
