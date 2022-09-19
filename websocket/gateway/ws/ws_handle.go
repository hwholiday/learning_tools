package ws

import (
	"fmt"
	"time"

	"github.com/hwholiday/learning_tools/websocket/pb"
)

func wsHandle(w *WsConnection) {
	var (
		err error
		msg interface{}
	)
	fmt.Println("开启链接: ", w.GetIp(), "UID", w.GetUid(), "WSID", w.GetWsId())
	for {
		if msg, err = w.ReadMsg(); err != nil {
			wsClose(w)
			return
		}
		switch msg.(type) {
		case *pb.Ping:
			fmt.Println("*pb.Ping", msg.(*pb.Ping))
			_ = w.ws.SetReadDeadline(time.Now().Add(time.Second * 20))
			_ = w.SendMsg(&pb.Pong{
				Times: time.Now().Unix(),
			})
		}
	}
}

func wsClose(w *WsConnection) {
	w.close()
}
