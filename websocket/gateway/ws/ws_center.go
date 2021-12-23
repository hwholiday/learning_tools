package ws

import (
	"sync"

	"github.com/hwholiday/learning_tools/websocket/gateway/msg"
)

var ctr *wsCenter

type wsCenter struct {
	allConn *sync.Map
	MaxConn int32 //允许的最大连接数
}

func NewWsCenter(maxConn int32) {
	ctr = &wsCenter{
		allConn: new(sync.Map),
		MaxConn: maxConn,
	}
}

func GetWsCenter() *wsCenter {
	return ctr
}

func (c *wsCenter) GetAllConn() *sync.Map {
	return c.allConn
}

func (c *wsCenter) AddConn(uid uint32, w *WsConnection) error {
	if c.CountConn()+1 > c.MaxConn {
		return msg.ErrExceedMaxConn
	}
	c.allConn.Store(uid, w)
	return nil
}

func (c *wsCenter) DeleteConn(uid int32) {
	c.allConn.Delete(uid)
}

func (c *wsCenter) CountConn() (count int32) {
	c.allConn.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return
}

func (c *wsCenter) Broadcast(data interface{}) {
	c.allConn.Range(func(key, value interface{}) bool {
		ws := value.(*WsConnection)
		_ = ws.SendMsg(data)
		return true
	})
}
