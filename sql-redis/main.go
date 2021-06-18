package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/groupcache/singleflight"
	"log"
	"time"
)

var errPlaceholder = errors.New("placeholder")
var errNotFind = errors.New("not find")

type cache struct {
	rds *redis.Client
	g   singleflight.Group
}

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr: "172.12.12.165:6379",
	})
	var c = &cache{
		rds: rds,
	}
	var str string
	err := c.Take(&str, "user:info:2", func(v interface{}) error {
		*v.(*string) = "我是你爸爸"
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func (c *cache) Take(v interface{}, key string, query func(v interface{}) error) error {
	return c.Take2(v, key, query, func(v interface{}) error {
		//保存DB数据到缓存
		fmt.Println("保存DB数据到缓存")
		return c.Set(key, v)
	})
}

func (c *cache) Take2(v interface{}, key string, query func(v interface{}) error, cacheVal func(v interface{}) error) error {
	val, err := c.g.Do(key, func() (interface{}, error) {
		//从缓存里面读取
		if err := c.Get(key, v); err != nil {
			if err == errPlaceholder {
				return nil, errNotFind
			} else if err != errNotFind {
				return nil, err
			}
			//err == errNotFind
			//从DB里面读取
			if err = query(v); err == errNotFind {
				if err = c.Set(key, "*"); err != nil {
					log.Println("set redis key  : ", key, " err :", err.Error())
				}
				return nil, errNotFind
			} else if err != nil {
				return nil, err
			}
			//调用保存缓存方法
			if err = cacheVal(v); err != nil {
				log.Println("cacheVal redis key  : ", key, " err :", err.Error())

			}
		}
		return json.Marshal(v)
	})
	if err != nil {
		return err
	}
	//等待到了 singleflight 的返回值并返回
	return json.Unmarshal(val.([]byte), v)
}

func (c *cache) Get(key string, v interface{}) error {
	data, err := c.rds.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return errNotFind
	}
	if err != nil {
		return err
	}
	if data == "*" {
		return errPlaceholder
	}
	if err = json.Unmarshal([]byte(data), v); err == nil {
		return nil
	}
	if err = c.rds.Del(key).Err(); err != nil {
		log.Println("del redis key  : ", key, " err :", err.Error())
	}
	return errNotFind
}

func (c *cache) Set(key string, v interface{}) error {
	fmt.Println("v", v)
	fmt.Println("key", key)
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.rds.Set(key, data, time.Second*1160).Err()
}
