package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type WsConnection struct {
	mu        sync.Mutex
	connId    string
	uid       uint32
	ws        *websocket.Conn
	readChan  chan []byte
	writeChan chan []byte
	closeChan chan bool
	isOpen    bool
	clientIp  string
}

var (
	WsErrConnLoss = errors.New("conn already close")
)

func NewWsConnection(conn *websocket.Conn) *WsConnection {
	ws := &WsConnection{}
	ws.ws = conn
	ws.readChan = make(chan []byte, 10)
	ws.writeChan = make(chan []byte, 10)
	ws.closeChan = make(chan bool)
	ws.isOpen = true
	ws.connId = uuid.NewV5(uuid.Must(uuid.NewV4()), "ws").String()
	go ws.read()
	go ws.send()
	return ws
}
func (w *WsConnection) SetIp(ip string) {
	w.clientIp = ip
}
func (w *WsConnection) GetIp() string {
	return w.clientIp
}

func (w *WsConnection) SetUid(uid uint32) {
	w.uid = uid
}
func (w *WsConnection) GetUid() uint32 {
	return w.uid
}

func (w *WsConnection) GetWsId() string {
	return w.connId
}

func (w *WsConnection) read() {
	var (
		Data []byte
		err  error
	)
	w.ws.SetReadLimit(2048)
	//_ = w.ws.SetReadDeadline(time.Now().Add(time.Second * 60))
	for {
		if _, Data, err = w.ws.ReadMessage(); err != nil {
			w.close()
			return
		}
		select {
		case w.readChan <- Data:
		case <-w.closeChan:
			return
		}
	}
}
func (w *WsConnection) send() {
	for {
		select {
		case message := <-w.writeChan:
			writer, err := w.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				w.close()
				return
			}
			_, _ = writer.Write(message)
			_ = writer.Close()
		case <-w.closeChan:
			return
		}
	}
}

func (w *WsConnection) ReadMsg() (message []byte, err error) {
	select {
	case message = <-w.readChan:
	case <-w.closeChan:
		err = WsErrConnLoss
	}
	return
}

func (w *WsConnection) SendMsg(msg []byte) (err error) {
	select {
	case w.writeChan <- msg:
	case <-w.closeChan:
		err = WsErrConnLoss
	}
	return
}

func (w *WsConnection) close() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.isOpen {
		fmt.Println("关闭链接: ", w.GetIp(), "ID", w.GetWsId())
		_ = w.ws.Close()
		w.isOpen = false
		w.closeChan <- true
	}
}
