package gateway

import "sync"

type PushInfo struct {
	Type   int
	roomId string
	info   []byte
}

type ConnManage struct {
     Room []*Room
     allConn sync.Map

}