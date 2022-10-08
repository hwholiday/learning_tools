package dove

import (
	"errors"
	"github.com/golang/protobuf/proto"
	api "github.com/hwholiday/ghost/dove/api/dove"
	"github.com/hwholiday/ghost/dove/network"
	"log"
	"net"
	"sync"
)

const (
	DefaultWsPort = ":8081"
)
const (
	DefaultConnAcceptCrcId uint64 = 1
	DefaultConnCloseCrcId  uint64 = 2
)

type HandleFunc func(cli network.Conn, reqData interface{})

type Dove interface {
	RegisterHandleFunc(id uint64, fn HandleFunc)
	Accept(conn net.Conn) error
}

type dove struct {
	rw            sync.RWMutex
	manger        *manager
	HandleFuncMap map[uint64]HandleFunc
}

func NewDove() Dove {
	h := dove{
		manger:        Manager(),
		HandleFuncMap: make(map[uint64]HandleFunc),
	}
	return &h
}

func (h *dove) verifyID(id uint64) error {
	if id == DefaultConnCloseCrcId || id == DefaultConnAcceptCrcId {
		return errors.New("please don't use default id")
	}
	return nil
}

func (h *dove) RegisterHandleFunc(id uint64, fn HandleFunc) {
	if err := h.verifyID(id); err != nil {
		log.Printf("[Dove] RegisterHandleFunc : %s \n", err.Error())
		return
	}
	h.rw.Lock()
	defer h.rw.Unlock()
	if _, ok := h.HandleFuncMap[id]; ok {
		log.Printf("[Dove] RegisterHandleFunc already register id : %d \n", id)
		return
	}
	h.HandleFuncMap[id] = fn
}

func (h *dove) triggerHandle(client network.Conn, id uint64, req []byte) {
	fn, ok := h.HandleFuncMap[id]
	if !ok {
		log.Printf("[Dove] accept HandleFuncMap not register id : %d \n", req)
		return
	}
	fn(client, req)
}

func (h *dove) Accept(conn net.Conn) error {
	client, err := network.NewConn(network.WithConn(conn))
	if err != nil {
		log.Printf("[Dove] Accept NewConn  %s \n", err.Error())
		return err
	}
	if err = h.manger.Add(client.Cache().Get(network.Identity).String(), client); err != nil {
		return err
	}
	h.triggerHandle(client, DefaultConnAcceptCrcId, nil)
	go func() {
		for {
			byt, err := client.Read()
			if err != nil {
				h.manger.Del(client.Cache().Get(network.Identity).String())
				h.triggerHandle(client, DefaultConnCloseCrcId, nil)
				log.Printf("[Dove] Accept Read  %s \n", err.Error())
				return
			}
			req, err := parseByt(byt)
			if err != nil {
				log.Printf("[Dove] Accept parseByt  %s \n", err.Error())
				continue
			}
			h.triggerHandle(client, req.GetHeader().GetCrcId(), req.GetBody().GetData())
		}
	}()
	return nil
}

func parseByt(byt []byte) (*api.Dove, error) {
	var req api.Dove
	if err := proto.Unmarshal(byt, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
