package gateway

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type WsConnection struct {
	mu        sync.Mutex
	connId    string
	ws        *websocket.Conn
	readChan  chan interface{}
	writeChan chan interface{}
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
	ws.readChan = make(chan interface{}, 10)
	ws.writeChan = make(chan interface{}, 10)
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

func (w *WsConnection) GetWsId() string {
	return w.connId
}

func (w *WsConnection) read() {
	var (
		Data []byte
		err  error
	)
	w.ws.SetReadLimit(1024)
	_ = w.ws.SetReadDeadline(time.Now().Add(time.Second * 20))
	for {
		if _, Data, err = w.ws.ReadMessage(); err != nil {
			w.close()
			return
		}
		var message interface{}
		if message, err = GetMsgProtocol().Unmarshal(Data); err != nil {
			w.close()
			return
		}
		select {
		case w.readChan <- message:
		case <-w.closeChan:
			return
		}
	}
}
func (w *WsConnection) send() {
	var (
		message interface{}
	)
	for {
		select {
		case message = <-w.writeChan:
			data, err := GetMsgProtocol().Marshal(message)
			if err != nil {
				w.close()
				return
			}
			writer, err := w.ws.NextWriter(websocket.BinaryMessage)
			if err != nil {
				w.close()
				return
			}
			_, _ = writer.Write(data)
			_ = writer.Close()
		case <-w.closeChan:
			return
		}
	}
}

func (w *WsConnection) ReadMsg() (message interface{}, err error) {
	select {
	case message = <-w.readChan:
	case <-w.closeChan:
		err = WsErrConnLoss
	}
	return
}

func (w *WsConnection) SendMsg(msg interface{}) (err error) {
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
