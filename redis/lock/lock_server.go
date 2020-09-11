package lock

type RedisLockServer interface {
	TryLock() error
	UnLock() error
	GetLockKey() string
	GetLockVal() string
}
