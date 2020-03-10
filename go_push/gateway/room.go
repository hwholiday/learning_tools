package gateway

import "sync"

//一个房间代表一个订阅推送类型
type Room struct {
	id    string
	title string
	RConn sync.Map
}

func NewRoom(id,title string) *Room {
	return &Room{
		id: id,
		title:title,
	}
}

func (r *Room) JoinRoom(ws *WsConnection) {
	if _, ok := r.RConn.Load(ws.connId); !ok {
		r.RConn.Store(ws.GetWsId(), ws)
	}
}

func (r *Room) LeaveRoom(ws *WsConnection) {
	if _, ok := r.RConn.Load(ws.GetWsId); ok {
		r.RConn.Delete(ws.GetWsId())
	}
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
