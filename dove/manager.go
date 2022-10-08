package dove

import (
	"errors"
	"github.com/hwholiday/ghost/dove/network"
	"sync"
	"sync/atomic"
)

var ErrExceedsLengthLimit = errors.New("exceeds length limit")

type manager struct {
	maxConn uint32
	connNum uint32
	connMap sync.Map
}

func Manager() *manager {
	return &manager{}
}

func (m *manager) Add(identity string, conn network.Conn) error {
	if m.GetConnNum() >= m.maxConn {
		return ErrExceedsLengthLimit
	}
	if old, ok := m.GetConn(identity); ok {
		//关闭老的链接信息，这里可能是异地登陆
		old.(network.Conn).Close()
		m.Del(identity)
	}
	m.connMap.Store(identity, conn)
	atomic.AddUint32(&m.connNum, 1)
	return nil
}

func (m *manager) Del(identity string) {
	if _, ok := m.connMap.Load(identity); !ok {
		return
	}
	atomic.AddUint32(&m.connNum, -1)
	m.connMap.Delete(identity)
}

func (m *manager) GetConnNum() uint32 {
	return atomic.LoadUint32(&m.connNum)
}

func (m *manager) GetConn(identity string) (network.Conn, bool) {
	val, ok := m.connMap.Load(identity)
	if !ok {
		return nil, false
	}
	return val.(network.Conn), true
}

func (m *manager) GetAllConn() []network.Conn {
	var clientArr = make([]network.Conn, 0, m.GetConnNum())
	m.connMap.Range(func(key, value any) bool {
		clientArr = append(clientArr, value.(network.Conn))
		return true
	})
	return clientArr
}
