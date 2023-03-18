package dove

import (
	"errors"
	"github.com/golang/protobuf/proto"
	api "github.com/hwholiday/ghost/dove/api/dove"
	"github.com/hwholiday/ghost/dove/network"
	"github.com/rs/zerolog/log"
	"sync"
)

type HandleFunc func(cli network.Conn, data *api.Dove)

type Dove interface {
	RegisterHandleFunc(id uint64, fn HandleFunc)
	Accept(opt ...network.Option) error
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
	setup()
	return &h
}

func (h *dove) RegisterHandleFunc(id uint64, fn HandleFunc) {
	h.rw.Lock()
	defer h.rw.Unlock()
	if _, ok := h.HandleFuncMap[id]; ok {
		log.Printf("[Dove] RegisterHandleFunc already register id : %d ", id)
		return
	}
	h.HandleFuncMap[id] = fn
}

func (h *dove) triggerHandle(client network.Conn, id uint64, data *api.Dove) {
	fn, ok := h.HandleFuncMap[id]
	if !ok {
		log.Printf("[Dove] accept HandleFuncMap not register id : %d ", id)
		return
	}
	fn(client, data)
}

func (h *dove) Accept(opt ...network.Option) error {
	client, err := network.NewConn(opt...)
	if err != nil {
		log.Printf("[Dove] Accept NewConn  %s ", err.Error())
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
				if !errors.Is(err, network.AlreadyCloseErr) {
					log.Printf("[Dove] Accept Read  %s ", err.Error())
				}
				return
			}
			req, err := parseByt(byt)
			if err != nil {
				log.Printf("[Dove] Accept parseByt  %s ", err.Error())
				continue
			}
			h.triggerHandle(client, req.GetMetadata().GetCrcId(), req)
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
