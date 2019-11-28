package lock

type RedisLockServer interface {
	TryLock() (bool, error)
	UnLock() error
	GetLockKey() string
	GetLockVal() string
}
