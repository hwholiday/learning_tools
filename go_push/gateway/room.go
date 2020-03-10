package gateway

import (
	"errors"
	"sync"
)

//一个房间代表一个订阅推送类型
type Room struct {
	id    int
	title string
	RConn sync.Map
}

func newRoom(id int, title string) *Room {
	return &Room{
		id:    id,
		title: title,
	}
}

func (r *Room) JoinRoom(ws *WsConnection) error {
	if _, ok := r.RConn.Load(ws.GetWsId()); ok {
		return errors.New("already exists")
	}
	r.RConn.Store(ws.GetWsId(), ws)
	return nil
}

func (r *Room) LeaveRoom(wsId string) error {
	if _, ok := r.RConn.Load(wsId); !ok {
		return errors.New("already delete")
	}
	r.RConn.Delete(wsId)
	return nil
}

func (r *Room) Count() int {
	var count int
	r.RConn.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

func (r *Room) Push(msg *WSMessage) {
	var (
		ws *WsConnection
		ok bool
	)
	r.RConn.Range(func(_, value interface{}) bool {
		if ws, ok = value.(*WsConnection); ok {
			_ = ws.SendMsg(msg)
		}
		return true
	})
}
