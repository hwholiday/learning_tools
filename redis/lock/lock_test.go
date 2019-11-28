package lock

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
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
	redisLock := NewRedisLock(GlobalClient, "1", "2", time.Second*3)
	InitRedis(redisLock)
	select {}
}
func InitRedis(lock RedisLockServer) {
	go func() {
		for {
			time.Sleep(time.Second)
			succ, err := lock.TryLock()
			if err != nil {
				panic(err)
			}
			if succ {
				fmt.Println("获取锁成功")
			} else {
				fmt.Println("获取锁失败")
			}
		}
	}()
}
