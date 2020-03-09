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

func (r *RoomManage) AddConn(ws *WsConnection) {
	r.AllConn.Store(ws.GetWsId(), ws)
}

func (r *RoomManage) DelConn(ws *WsConnection) {
	r.AllConn.Delete(ws.GetWsId())
}

func (r *RoomManage) AddRoom(id string, ws *WsConnection) error {
	val, ok := r.AllRoom.Load(id)
	if !ok {
		//room = NewRoom(id)
		//r.AllRoom.Store(id, room)
		return errors.New("not find room")
	}
	room := val.(*Room)
	room.JoinRoom(ws)
	return nil
}

func (r *RoomManage) LeaveRoom(id string, ws *WsConnection) error {
	val, ok := r.AllRoom.Load(id)
	if !ok {
		return errors.New("not find room")
	}
	room := val.(*Room)
	room.LeaveRoom(ws)
	if room.Count() <= 0 {
		r.AllRoom.Delete(room.id)
	}
	return nil
}

func (r *RoomManage) PushRoom(id string, msg *WSMessage) error {
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
