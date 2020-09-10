package gcache

import "container/list"

type LRU struct {
	maxByte  int64
	useByte  int64
	ll       *list.List
	cache    map[string]*list.Element
	CallBack func(key string, val []byte)
}

type Node struct {
	Key string
	Val []byte
}

func NewLRU(maxByte int64, callBack func(key string, val []byte)) *LRU {
	return &LRU{
		maxByte:  maxByte,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		CallBack: callBack,
	}
}

func (g *LRU) Add(key string, data []byte) {
	if val, ok := g.cache[key]; ok {
		g.ll.MoveToFront(val)
	} else {
		g.useByte += int64(len(key)) + int64(len(data))
		val := g.ll.PushFront(&Node{
			Key: key,
			Val: data,
		})
		g.cache[key] = val

	}
	for g.maxByte != 0 && g.maxByte < g.useByte {
		g.Remove()
	}
}

func (g *LRU) Del(key string) {
	g.Remove(key)
}

func (g *LRU) Remove(k ...string) {
	var (
		val *list.Element
		ok  bool
	)
	if len(k) == 0 {
		val = g.ll.Back()
	} else {
		if val, ok = g.cache[k[0]]; !ok {
			return
		}
	}
	if val != nil {
		g.ll.Remove(val)
		node := val.Value.(*Node)
		delete(g.cache, node.Key)
		g.useByte -= int64(len(node.Key)) + int64(len(node.Val))
		if g.CallBack != nil {
			g.CallBack(node.Key, node.Val)
		}
	}
}

func (g *LRU) Get(key string) (v []byte, ok bool) {
	var val *list.Element
	if val, ok = g.cache[key]; ok {
		g.ll.MoveToFront(val)
		node := val.Value.(*Node)
		v = node.Val
		return
	}
	return
}

func (g *LRU) Len() int {
	return g.ll.Len()
}
