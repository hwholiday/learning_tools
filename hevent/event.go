package hevent

import "sync"

type HEvent struct {
	Data  interface{}
	Topic string
}

type HEventData chan HEvent
type HEventDataArray []HEventData //一个topic 可以有多个消费者

type HEventBus struct {
	sub map[string]HEventDataArray
	rm  sync.RWMutex
}

var h *HEventBus

func init() {
	h = &HEventBus{
		sub: make(map[string]HEventDataArray),
	}
}

type HEventRepo interface {
	Sub(topic string, ch HEventData)
	Push(topic string, data interface{})
}

func HEventSrv() *HEventBus {
	return h
}

func (h *HEventBus) Sub(topic string, ch HEventData) {
	h.rm.Lock()
	if chanEvent, ok := h.sub[topic]; ok {
		h.sub[topic] = append(chanEvent, ch)
	} else {
		h.sub[topic] = append([]HEventData{}, ch)
	}
	defer h.rm.Unlock()
}

func (h *HEventBus) Push(topic string, data interface{}) {
	h.rm.RLock()
	defer h.rm.RUnlock()
	if chanEvent, ok := h.sub[topic]; ok {
		for _, ch := range chanEvent {
			ch <- HEvent{
				Data:  data,
				Topic: topic,
			}
		}
	}
}
