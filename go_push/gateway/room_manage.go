package gateway

import (
	"errors"
	"sync"
)

var (
	manage *RoomManage
)

type RoomManage struct {
	AllRoom sync.Map
	AllConn sync.Map
}

func NewRoomManage() {
	manage = &RoomManage{}
	return
}
func GetRoomManage() *RoomManage {
	return manage
}

func (r *RoomManage) NewRoom(id int, title string) error {
	_, ok := r.AllRoom.Load(id)
	if ok {
		return errors.New("already exists")
	}
	r.AllRoom.Store(id, newRoom(id, title))
	return nil
}

func (r *RoomManage) AddConn(ws *WsConnection) {
	r.AllConn.Store(ws.GetWsId(), ws)
}

func (r *RoomManage) DelConn(ws *WsConnection) {
	r.AllConn.Delete(ws.GetWsId())
}

func (r *RoomManage) AddRoom(id int, wsId string) error {
	var room *Room
	var ws *WsConnection
	val, ok := r.AllRoom.Load(id)
	if !ok {
		return errors.New("not find room")
	}
	wsVal, ok := r.AllConn.Load(wsId)
	if !ok {
		return errors.New("not find conn")
	}
	room = val.(*Room)
	ws = wsVal.(*WsConnection)
	if err := room.JoinRoom(ws); err != nil {
		return err
	}
	ws.addRoom.Store(id, true)
	return nil
}

func (r *RoomManage) LeaveRoom(id int, wsId string) error {
	var room *Room
	val, ok := r.AllRoom.Load(id)
	if !ok {
		return errors.New("not find room")
	}
	room = val.(*Room)
	if err := room.LeaveRoom(wsId); err != nil {
		return err
	}
	/*if room.Count() <= 0 {
	    //房间人数为0 删除房间
		r.AllRoom.Delete(room.id)
	}*/
	return nil
}

func (r *RoomManage) PushRoom(id int, msg *WSMessage) error {
	val, ok := r.AllRoom.Load(id)
	if !ok {
		return errors.New("not find room")
	}
	room := val.(*Room)
	room.Push(msg)
	return nil
}

func (r *RoomManage) PushAll(msg *WSMessage) {
	r.AllConn.Range(func(_, value interface{}) bool {
		if ws, ok := value.(*WsConnection); ok {
			_ = ws.SendMsg(msg)
		}
		return true
	})
}
