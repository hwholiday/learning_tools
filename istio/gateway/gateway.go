package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strings"
	"time"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:   2048,
	WriteBufferSize:  2048,
	HandshakeTimeout: 5 * time.Second,
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/connect", getConn)
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("test gateway success"))
	})
	// HTTP服务
	server := http.Server{
		Addr:         ":8888",
		ReadTimeout:  time.Duration(20) * time.Second,
		WriteTimeout: time.Duration(20) * time.Second,
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
