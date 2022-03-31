package gateway

import (
	"fmt"
	"time"
)

func (w *WsConnection) WsHandle() {
	var (
		err error
		msg *WSMessage
	)
	fmt.Println("开启链接: ", w.GetIp(), "ID", w.GetWsId())
	GetRoomManage().AddConn(w)
	for {
		if msg, err = w.ReadMsg(); err != nil {
			w.CloseConn()
			return
		}
		//http://www.easyswoole.com/wstool.html  测试地址
		switch {
		case msg.Type == 1: //{"type":1,"data":"PING"},{"type":1,"data":"PONG"}
			_ = w.ws.SetReadDeadline(time.Now().Add(time.Second * 10))
			_ = w.SendMsg(&WSMessage{Type: 1, Data: "PONG"})
		default:
			fmt.Println("OTHER", msg.Type, msg.Data)

		}
	}

}

func (w *WsConnection) CloseConn() {
	w.close()
	GetRoomManage().DelConn(w)
	w.addRoom.Range(func(key, _ interface{}) bool {
		_ = GetRoomManage().LeaveRoom(key.(int), w.GetWsId())
		w.addRoom.Delete(key)
		return true
	})
}
