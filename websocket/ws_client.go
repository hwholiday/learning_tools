package wshandel

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net"
	"net/http"
	"sync"
	"time"
)

var WsClientUpgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, ReadBufferSize: 4096,
	WriteBufferSize:   1024,
}

type WsClient struct {
	conn         *websocket.Conn
	connId       string
	isConn       bool
	tag          string
	deviceId     int
	remote       string
	Token        string
	isLogin      bool
	mutex        *sync.Mutex
	waitGroup    *sync.WaitGroup
	channelRead  chan []byte
	channelWrite chan []byte
}

func NewWsClient(ws *websocket.Conn) *WsClient {
	return &WsClient{
		conn:      ws,
		isLogin:   false,
		isConn:    true,
		waitGroup: new(sync.WaitGroup),
		mutex:     new(sync.Mutex),
		connId:    uuid.NewV5(uuid.Must(uuid.NewV4()), "ws4567").String(),
	}
}

func (w *WsClient) SetTag(tag string) {
	w.tag = tag
}

func (w *WsClient) GetTag() string {
	return w.tag
}
func (w *WsClient) SetDeviceId(deviceId int) {
	w.deviceId = deviceId
}

func (w *WsClient) GetDeviceId() int {
	return w.deviceId
}

func (w *WsClient) SetRemote(ip string) {
	w.remote = ip
}

func (w *WsClient) GetRemote() string {
	return w.remote
}

func (w *WsClient) GetConn() *websocket.Conn {
	return w.conn
}

func (w *WsClient) GetConnId() string {
	return w.connId
}

func (w *WsClient) OpenChannel() {
	w.channelRead = make(chan []byte, 10)
	w.channelWrite = make(chan []byte, 10)
	w.waitGroup.Add(2)
	go w.readMessage()
	go w.writeMessage()
}

func (w *WsClient) CloseChannel() {
	close(w.channelRead)
	close(w.channelWrite)
}

func (w *WsClient) Close() error {
	w.mutex.Lock()
	w.isConn = false
	w.mutex.Unlock()
	return w.conn.Close()
}

func (w *WsClient) LocalAddr() net.Addr {
	return w.conn.LocalAddr()
}

func (w *WsClient) RemoteAddr() net.Addr {
	return w.conn.RemoteAddr()
}

func (b *WsClient) SendMessage(msg []byte) {
	if b.isConn {
		b.channelWrite <- msg
	} else {
		fmt.Println("链接已经关闭")
	}
}

//从连接读取数据,并写入channel,用于分包
func (b *WsClient) ReadBytesFromConn() error {
	b.conn.SetReadLimit(4096)
	for {
		_ = b.conn.SetReadDeadline(time.Now().Add(40 * time.Second))
		_, p, err := b.conn.ReadMessage()
		if err != nil {
			return err
		}
		b.channelRead <- p
	}
}

func (b *WsClient) ClientClose() {
	if b.isConn {
		_ = b.conn.WriteMessage(websocket.CloseMessage, []byte{})
		_ = b.Close()
		//RemoveSession(b.GetTag()) 移除会话session
		b.CloseChannel()
	}
	b.waitGroup.Wait()
}

func (w *WsClient) readMessage() {
	for {
		if !w.isConn {
			w.waitGroup.Done()
			return
		}
		select {
		case msg, ok := <-w.channelRead:
			if !ok || !w.isConn {
				w.waitGroup.Done()
				return
			}
			fmt.Println("上传参数", string(msg))
			//ProtocolAnalysis(msg, w) 解析分析
		}
	}
}

func (w *WsClient) writeMessage() {
	for {
		if !w.isConn {
			w.waitGroup.Done()
			return
		}
		select {
		case msg, ok := <-w.channelWrite:
			if !ok || !w.isConn {
				w.waitGroup.Done()
				return
			}
			if err := w.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				return
			}
		}
	}
}
