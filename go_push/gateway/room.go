package gateway

import "sync"

//一个房间代表一个订阅推送类型
type Room struct {
	mu    sync.Mutex
	title string
	index int
	RConn sync.Map
}

func NewRoom(id int, title string) *Room {
	return &Room{
		title: title,
		index: id,
	}
}

func (r *Room) JoinRoom(ws *WsConnection) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.RConn.Load(ws.connId); !ok {
		r.RConn.Store(ws.GetWsId(), ws)
	}
}

func (r *Room) LeaveRoom(ws *WsConnection) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.RConn.Load(ws.connId); ok {
		r.RConn.Delete(ws.GetWsId())
	}
}

func (r *Room) Push(msg *WSMessage) {
	var (
		ws *WsConnection
		ok bool
	)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.RConn.Range(func(_, value interface{}) bool {
		if ws, ok = value.(*WsConnection); ok {
			_ = ws.SendMsg(msg)
		}
		return true
	})
}
