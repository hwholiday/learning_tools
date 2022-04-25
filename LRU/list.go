package LRU

import (
	"container/list"
	"errors"
	"sync"
)

type CallBack func(key interface{}, value interface{})

type Lru struct {
	max   int
	list  *list.List
	Call  CallBack
	cache map[interface{}]*list.Element
	mu    *sync.RWMutex
}

type Node struct {
	Key interface{}
	Val interface{}
}

func NewLru(len int, c CallBack) *Lru {
	return &Lru{
		max:   len,
		list:  list.New(),
		Call:  c,
		cache: make(map[interface{}]*list.Element),
		mu:    new(sync.RWMutex),
	}
}

func (l *Lru) Add(key interface{}, val interface{}) error {
	if l.list == nil {
		return errors.New("not init NewLru")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.cache[key]; ok { //以及存在
		e.Value.(*Node).Val = val
		l.list.MoveToFront(e)
		return nil
	}
	ele := l.list.PushFront(&Node{
		Key: key,
		Val: val,
	})
	l.cache[key] = ele
	if l.max != 0 && l.list.Len() > l.max {
		if e := l.list.Back(); e != nil {
			l.list.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache, node.Key)
			if l.Call != nil {
				l.Call(node.Key, node.Val)
			}
		}
	}
	return nil
}

func (l *Lru) Get(key interface{}) (val interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	l.mu.RLock()
	defer l.mu.RUnlock()
	if ele, ok := l.cache[key]; ok {
		l.list.MoveToFront(ele)
		return ele.Value.(*Node).Val, true
	}
	return
}

func (l *Lru) GetAll() []*Node {
	l.mu.RLock()
	defer l.mu.RUnlock()
	var data []*Node
	for _, v := range l.cache {
		data = append(data, v.Value.(*Node))
	}
	return data
}

func (l *Lru) Del(key interface{}) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if ele, ok := l.cache[key]; ok {
		delete(l.cache, ele)
		if e := l.list.Back(); e != nil {
			l.list.Remove(e)
			delete(l.cache, key)
			if l.Call != nil {
				node := e.Value.(*Node)
				l.Call(node.Key, node.Val)
			}
		}
	}
}
