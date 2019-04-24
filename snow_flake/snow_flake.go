package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

// 起始时间戳 (毫秒数显示)
var Epoch int64 = 1288834974657 // timestamp 2006-03-21:20:50:14 GMT

// ID 结构
type ID int64

// 存储基础信息的 Node 结构
type Node struct {
	mu        sync.Mutex // 保证并发安全
	timestamp int64
	node      int64
	step      int64
}

// 生成、返回唯一 snowflake ID
func (n *Node) Generate() ID {
	n.mu.Lock()         // 保证并发安全, 加锁
	defer n.mu.Unlock() // 解锁
	// 获取当前时间的时间戳 (毫秒数显示)
	now := time.Now().UnixNano() / 1e6
	if n.timestamp == now {
		// step 步进 1
		n.step ++
		// 当前 step 用完
		if n.step > stepMax {
			// 等待本毫秒结束
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 本毫秒内 step 用完
		n.step = 0
	}
	n.timestamp = now
	result := ID((now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step))
	return result
}
// 生成一个新的 Node 类型变量
func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}
func main() {
	// 生成节点实例
	node, err := NewNode(1)
	if err!=nil{
		panic(err)
	}
	for  {
		fmt.Println(node.Generate())
	}
}