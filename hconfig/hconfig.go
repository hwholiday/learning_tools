package hconfig

import (
	"context"
	"errors"
	"github.com/hwholiday/learning_tools/hconfig/hconf"
	"sync"
)

type config struct {
	opts    *options
	cache   *sync.Map
	watcher hconf.DataWatcher
}
type HConfig interface {
	Load() error
	Get(key string) (HVal, error)
	Watch(event WatchEvent) error
	Close() error
}

type WatchEvent func(path string, v HVal)

func NewHConfig(opts ...Option) (HConfig, error) {
	options, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}
	return &config{
		opts:  options,
		cache: new(sync.Map),
	}, nil
}

func (c *config) Load() error {
	kvs, err := c.opts.dataSource.Load()
	if err != nil {
		return nil
	}
	for _, v := range kvs {
		c.cache.Store(v.Key, v.Val)
	}
	return nil
}

func (c *config) Get(key string) (HVal, error) {
	val, ok := c.cache.Load(key)
	if !ok {
		return nil, errors.New("not find")
	}
	return HVal(val.([]byte)), nil
}

func (c *config) Watch(event WatchEvent) error {
	var err error
	if c.watcher, err = c.opts.dataSource.Watch(); err != nil {
		return err
	}
	go c.watch(event)
	return nil
}

func (c *config) watch(event WatchEvent) {
	for {
		kvs, err := c.watcher.Change()
		if errors.Is(err, context.Canceled) {
			return
		}
		if err != nil {
			continue
		}
		for _, v := range kvs {
			c.cache.Store(v.Key, v.Val)
			event(v.Key, HVal(v.Val))
		}
	}
}

func (c *config) Close() error {
	c.cache = new(sync.Map)
	if c.watcher != nil {
		return c.watcher.Close()
	}
	return nil
}
