package lock

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisLock struct {
	conn    *redis.Client
	timeout time.Duration
	key     string
	val     string
}

func NewRedisLock(conn *redis.Client, key, val string, timeout time.Duration) *RedisLock {
	return &RedisLock{conn: conn, timeout: timeout, key: key, val: val}
}

// return nil ===> Get the lock successfully
func (lock *RedisLock) TryLock() error {
	ok, err := lock.conn.SetNX(lock.key, lock.val, lock.timeout).Result()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("try lock %s failed", lock.key)
	}
	return nil
}

func (lock *RedisLock) UnLock() error {
	luaDel := redis.NewScript("if redis.call('get',KEYS[1]) == ARGV[1] then " +
		"return redis.call('del',KEYS[1]) else return 0 end")
	return luaDel.Run(context.Background(), lock.conn, []string{lock.key}, lock.val).Err()
}

func (lock *RedisLock) GetLockKey() string {
	return lock.key
}

func (lock *RedisLock) GetLockVal() string {
	return lock.val
}
