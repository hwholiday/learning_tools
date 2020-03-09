package gateway

import "fmt"

func (w *WsConnection) WsHandle() {
	var (
		err error
		msg *WSMessage
	)
	GetRoomManage().AddConn(w)
	for {
		if msg, err = w.ReadMsg(); err != nil {
			w.CloseConn()
		}
		fmt.Println(msg)
	}

}

func (w *WsConnection) CloseConn() {
	w.close()
	GetRoomManage().DelConn(w)
	w.addRoom.Range(func(key, _ interface{}) bool {
		_ = GetRoomManage().LeaveRoom(key.(string), w)
		w.addRoom.Delete(key)
		return true
	})
}
