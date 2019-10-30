package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "",
			PoolSize:     10,
			DB:           0,
		},
	)
	err := GlobalClient.Ping().Err()
	if nil != err {
		panic(err)
	}
	//redis乐观锁支持，可以通过watch监听一些Key, 如果这些key的值没有被其他人改变的话，才可以提交事务。
	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get("pipe_test").Result()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		fmt.Println(v)
		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set("pipe_test", "new value 1111111111111", 0)
			return nil
		})
		return err
	}
	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
	err = GlobalClient.Watch(fn, "pipe_test")
	if nil != err {
		panic(err)
	}
}
