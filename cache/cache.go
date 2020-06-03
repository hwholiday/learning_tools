package cache

import (
	"sort"
	"sync"
	"time"
	"unsafe"
)

type cache struct {
	maxCache   int
	useCache   int
	mu         sync.RWMutex
	gcInterval time.Duration
	objects    map[string]Object
}

func NewCache(gc time.Duration) *cache {
	c := &cache{
		gcInterval: gc,
		objects:    make(map[string]Object),
	}
	go c.runGc()
	return c
}

func (c *cache) SetMaxMemory(size string) bool {
	num, err := GetSize(size)
	if err != nil {
		return false
	}
	c.mu.Lock()
	c.maxCache = num
	c.mu.Unlock()
	return true
}

func (c *cache) Set(key string, val interface{}, expire time.Duration) {
	var expiration int64
	if expire > 0 {
		expiration = time.Now().Add(expire).UnixNano()
	}
	valSize := int(unsafe.Sizeof(val))
	if c.useCache+valSize > c.maxCache {
		// 1 直接返回错误不让添加
		//log.Fatal(ErrNotEnoughSpace)
		//return
		// 2 处理掉使用次数最低的项目,达到需要存入的内存大小
		c.ClearSize(valSize)
	}
	c.mu.Lock()
	c.objects[key] = Object{
		Data:       val,
		Size:       valSize,
		Expiration: expiration,
	}
	c.maxCache += valSize
	c.mu.Unlock()
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	obj, ok := c.objects[key]
	if !ok {
		c.mu.Unlock()
		return nil, false
	}
	obj.Inquire++
	c.objects[key] = obj
	if obj.Expiration > 0 {
		if time.Now().UnixNano() > obj.Expiration {
			c.mu.RUnlock()
			return nil, false
		}
	}
	c.mu.RUnlock()
	return obj.Data, true
}

func (c *cache) Exists(key string) bool {
	c.mu.RLock()
	obj, ok := c.objects[key]
	if !ok {
		c.mu.Unlock()
		return false
	}
	obj.Inquire++
	c.objects[key] = obj
	if obj.Expiration > 0 {
		if time.Now().UnixNano() > obj.Expiration {
			c.mu.RUnlock()
			return false
		}
	}
	c.mu.RUnlock()
	return true
}

func (c *cache) Del(key string) bool {
	delete(c.objects, key)
	return true
}

func (c *cache) Flush() bool {
	c.mu.Lock()
	c.objects = map[string]Object{}
	c.mu.Unlock()
	return true
}

func (c *cache) Keys() int64 {
	var l int64
	c.mu.RLock()
	l = int64(len(c.objects))
	c.mu.Unlock()
	return l
}

func (c *cache) runGc() {
	ticker := time.NewTicker(c.gcInterval)
	for {
		select {
		case <-ticker.C:
			c.clear()
		}
	}
}

// 清理内存
func (c *cache) clear() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	for key, val := range c.objects {
		if val.Expiration > 0 && val.Expiration < now {
			c.Del(key)
		}
	}
	c.mu.Unlock()
}

// 清理内存
func (c *cache) ClearSize(need int) {
	var (
		nodes   []node
		needDel []string
		allSize int
	)
	c.mu.Lock()
	for key, val := range c.objects {
		nodes = append(nodes, node{size: val.Size, key: key, num: val.Inquire})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].num < nodes[j].num
	})
	for _, v := range nodes {
		if allSize > need {
			break
		}
		allSize += v.size
		needDel = append(needDel, v.key)
	}
	for _, v := range needDel {
		delete(c.objects, v)
	}
	c.mu.Unlock()
}
