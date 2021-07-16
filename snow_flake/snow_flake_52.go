package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	numberBits52 uint8 = 11
	numberMax52  int64 = -1 ^ (-1 << numberBits52)
	timeShift52  uint8 = 11
	startTime52  int64 = 1626408485000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type SnowWorker struct {
	mu        sync.Mutex
	timestamp int64
	number    int64
}

// NewSnowWorker 一毫秒发2047个
// 41位时间戳 11位自旋
func NewSnowWorker() *SnowWorker {
	// 生成一个新节点
	return &SnowWorker{
		timestamp: 0,
		number:    0,
	}
}

func (w *SnowWorker) GetTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *SnowWorker) GetId() (int64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := w.GetTime()
	if w.timestamp == now {
		w.number++
		if w.number > numberMax52 {
			for now <= w.timestamp {
				now = w.GetTime()
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	if now < w.timestamp {
		return 0, errors.New("clock move backwards")
	}
	ID := int64((now-startTime52)<<timeShift52 | (w.number))
	return ID, nil
}
func main() {
	// 生成节点实例
	node := NewSnowWorker()
	for {
		fmt.Println(node.GetId())
		return
	}
}
