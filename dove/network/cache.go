package network

import "sync"

type Cache struct {
	mp  sync.Map
	tmp any
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Save(k string, v any) {
	c.mp.Store(k, v)
}

func (c *Cache) Get(k string) *Cache {
	c.tmp = k
	return c
}
func (c *Cache) load() (any, bool) {
	defer c.clearTmp()
	return c.mp.Load(c.tmp)
}

func (c *Cache) Int() int {
	v, ok := c.load()
	if !ok {
		return 0
	}
	vStr, ok := v.(int)
	if !ok {
		return 0
	}
	return vStr
}
func (c *Cache) clearTmp() {
	c.tmp = nil
}

func (c *Cache) Result() (any, bool) {
	return c.load()
}

func (c *Cache) String() string {
	v, ok := c.load()
	if !ok {
		return ""
	}
	vStr, ok := v.(string)
	if !ok {
		return ""
	}
	return vStr
}
