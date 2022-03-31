package handler

import (
	"errors"
	"log"
	"strconv"
	"sync"
)

type hb struct {
	// 注册了的连接器
	connections sync.Map

	// 从连接器中发入的信息
	broadcast chan *ClientsReport

	// 从连接器中注册请求
	register chan *connection

	// 从连接器中注销请求
	unregister chan *connection

	//用户对照
	user sync.Map
}

var H = hb{
	broadcast:  make(chan *ClientsReport),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *hb) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections.Store(c, c.uid)
			h.user.Store(c.uid, c)
			c.send <- []byte("ok")
		case c := <-h.unregister:
			if _, ok := h.connections.Load(c); ok {
				h.user.Delete(c.uid)
				h.connections.Delete(c)
				close(c.send)
			}
		case m := <-h.broadcast:
			PushMsg(m.Uid, m.Msg)
		}
	}
}

func PushMsg(uid int, msg string) error {
	conn, ok := H.user.Load(uid)
	if !ok {
		return errors.New("该用户未在服务器登陆")
	}
	sendConn, ok := conn.(*connection)
	if !ok {
		return errors.New("获取该用户信息失败")
	}
	select {
	case sendConn.send <- []byte(msg):
		return nil
	default:
		H.user.Delete(uid)
		H.connections.Delete(conn)
		close(sendConn.send)
		log.Println("server_68:删除缓存信息:" + strconv.Itoa(uid))
		return errors.New("删除缓存信息")
	}

}
