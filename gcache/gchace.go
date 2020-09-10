package gcache

import (
	"fmt"
	"sync"
)

type Group struct {
	name   string
	getter Getter
	cache  cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

func NewGroups(name string, cacheBytes int64, getter Getter) *Group {
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		cache:  cache{maxCacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, fmt.Errorf("key is empty")
	}
	if val, ok := g.cache.get(key); ok {
		return val, nil
	}
	return g.load(key)
}

//读取其他地方的数据并填入缓存
func (g *Group) load(key string) ([]byte, error) {
	data, err := g.getter.Get(key)
	if err != nil {
		return nil, err
	}
	g.cache.add(key, data)
	return data, nil
}
