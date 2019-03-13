package wshandel

import (
	"github.com/gin-gonic/gin"
	"github.com/hwholiday/libs/logtool"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func NativityGetIp(r *http.Request) (data string) {
	if len(r.Header.Get("x-forwarded-for")) <= 0 {
		data = strings.Split(r.RemoteAddr, ":")[0]
	} else {
		data = strings.Split(r.Header.Get("x-forwarded-for"), ",")[0]
	}
	return data
}

//ws的入口
func ServeWs(g *gin.Context) {
	ip := NativityGetIp(g.Request)
	conn, err := WsClientUpgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		logtool.Zap.Error("ServeWs", zap.Error(err))
		return
	}
	wsConn := NewWsClient(conn)
	defer func() {
		if r := recover(); r != nil {
			wsConn.ClientClose()
		}
	}()
	wsConn.SetRemote(ip)
	//打开读写channel
	wsConn.OpenChannel()
	wsConn.conn.SetPingHandler(nil)
	err = wsConn.ReadBytesFromConn()
	//读取到错误或者EOF
	if err != nil {
		//正常的关闭
		logtool.Zap.Debug("ServeWs stopped because of: ", zap.String("IP", wsConn.remote), zap.Error(err))
		wsConn.ClientClose()
	}
}
