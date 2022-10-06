package dove

import (
	"github.com/hwholiday/ghost/dove/network"
	"log"
	"net"
	"sync"
)

const (
	DefaultWsPort = ":8081"
)
const (
	DefaultConnAcceptCrcId int = 1
	DefaultConnCloseCrcId  int = 2
)

type HandleFunc func(cli network.Conn, reqData interface{})

type Dove interface {
	RegisterHandleFunc(id int, fn HandleFunc)
	ListenWs() error
}

type dove struct {
	rw            sync.RWMutex
	manger        *manager
	HandleFuncMap map[int]HandleFunc
}

func NewDove() Dove {
	h := dove{
		manger:        Manager(),
		HandleFuncMap: make(map[int]HandleFunc),
	}
	return &h
}

func (h *dove) RegisterHandleFunc(id int, fn HandleFunc) {
	h.rw.Lock()
	defer h.rw.Unlock()
	if _, ok := h.HandleFuncMap[id]; ok {
		log.Printf("[Dove] RegisterHandleFunc already register id : %d \n", id)
		return
	}
	h.HandleFuncMap[id] = fn
}

func (h *dove) ListenWs() error {
	return nil
}

func (h *dove) accept(conn net.Conn) error {
	client, err := network.NewConn(network.WithConn(conn))
	if err != nil {
		return err
	}
	if err = h.manger.Add(client.Cache().Get(network.Identity).String(), client); err != nil {
		return err
	}
	for {
		byt, err := client.Read()
		if err != nil {
			h.manger.Del(client.Cache().Get(network.Identity).String())
			return err
		}
		fn, ok := h.HandleFuncMap[1]
		if !ok {
			log.Printf("[Dove] HandleFunc not register id : %d \n")
			continue
		}
		fn(client, byt)
	}
}
