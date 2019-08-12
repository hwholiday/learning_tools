package tool

import (
	"github.com/go-redis/redis"
	"time"
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

//return true ===> Get the lock successfully
func (lock *RedisLock) TryLock() (bool, error) {
	return lock.conn.SetNX(lock.key, lock.val, lock.timeout).Result()
}

func (lock *RedisLock) UnLock() error {
	return lock.conn.Del(lock.key).Err()
}

func (lock *RedisLock) GetLockKey() string {
	return lock.key
}

func (lock *RedisLock) GetLockVal() string {
	return lock.val
}
