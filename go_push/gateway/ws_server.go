package gateway

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strings"
	"time"
)

var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var (
	RoomTitle = []string{"健身", "体育", "电影", "音乐"}
)

func InitWsServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/connect", getConn)
	NewRoomManage()
	//初始化房间
	for i := 0; i < 4; i++ {
		_ = GetRoomManage().NewRoom(i, RoomTitle[i])
	}
	NewPushTask(len(RoomTitle), 3, 10)
	// HTTP服务
	server := http.Server{
		Addr:         "0.0.0.0:8888",
		ReadTimeout:  time.Duration(10) * time.Millisecond,
		WriteTimeout: time.Duration(10) * time.Millisecond,
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
	ws.WsHandle()
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
