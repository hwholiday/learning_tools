package ws

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hwholiday/learning_tools/websocket/gateway/msg"
	"github.com/hwholiday/learning_tools/websocket/pb"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWsServer() {
	msg.NewMsgProtocol(true)
	msg.GetMsgProtocol().Register(&pb.Ping{}, 1)
	msg.GetMsgProtocol().Register(&pb.Pong{}, 2)
	NewWsCenter(10000)
	mux := http.NewServeMux()
	mux.HandleFunc("/connect", getConn)
	// HTTP服务
	server := http.Server{
		Addr:         "0.0.0.0:8888",
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
		Handler:      mux,
	}
	fmt.Println("启动WS服务器成功 ：", 8888)
	_ = server.ListenAndServe()

}

func getConn(res http.ResponseWriter, req *http.Request) {
	var (
		err    error
		wsConn *websocket.Conn
	)
	if wsConn, err = wsUpgrader.Upgrade(res, req, nil); err != nil {
		return
	}
	ws := NewWsConnection(wsConn)
	ws.SetIp(ClientIP(req))
	ws.SetUid(uint32(time.Now().Unix()))
	wsHandle(ws)
}

func ClientIP(c *http.Request) string {
	clientIP := c.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(c.Header.Get(("X-Real-Ip")))
	}
	if clientIP != "" {
		return clientIP
	}
	addr := c.Header.Get("X-Appengine-Remote-Addr")
	if addr != "" {
		return addr
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
