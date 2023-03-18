package main

import (
	"github.com/gorilla/websocket"
	"github.com/hwholiday/ghost/dove"
	api "github.com/hwholiday/ghost/dove/api/dove"
	"github.com/hwholiday/ghost/dove/network"
	"github.com/rs/zerolog/log"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var client dove.Dove

func main() {
	dove.SetMode(dove.DebugMode)
	dove.SetConnMax(100)
	client = dove.NewDove()
	client.RegisterHandleFunc(dove.DefaultConnAcceptCrcId, func(cli network.Conn, data *api.Dove) {
		log.Info().Str("Identity", cli.Cache().Get(network.Identity).String()).Msg("设备上线")
	})
	client.RegisterHandleFunc(dove.DefaultConnCloseCrcId, func(cli network.Conn, data *api.Dove) {
		log.Info().Str("Identity", cli.Cache().Get(network.Identity).String()).Msg("设备离线")
	})
	Listen()
}

func Listen() {
	http.HandleFunc("/socket.io", getWsConn)
	log.Info().Str("addr", dove.DefaultWsPort).Msg("example service start succeed")
	err := http.ListenAndServe(dove.DefaultWsPort, nil)
	if err != nil {
		panic(err)
	}
}
func getWsConn(res http.ResponseWriter, req *http.Request) {
	var (
		err    error
		wsConn *websocket.Conn
	)
	if wsConn, err = wsUpgrader.Upgrade(res, req, nil); err != nil {
		return
	}
	go func() {
		err = client.Accept(network.WithConn(wsConn.UnderlyingConn()))
		if err != nil {
			log.Error().Err(err).Msg("Accept failed")
		}
	}()
}
