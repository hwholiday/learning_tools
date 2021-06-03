package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

type EtcdLock struct {
	client  *clientv3.Client
	timeout int64
	ctx     context.Context
	cancel  context.CancelFunc
	key     string
	val     string
	mutex   *concurrency.Mutex
	session *concurrency.Session
}

func NewEtcdLock(conn *clientv3.Client, key string, timeout int64) *EtcdLock {
	return &EtcdLock{client: conn, timeout: timeout, key: key}
}

func (lock *EtcdLock) TryLock() error {
	lock.ctx, lock.cancel = context.WithTimeout(context.Background(), time.Duration(lock.timeout)*time.Second)
	response, err := lock.client.Grant(lock.ctx, lock.timeout)
	if err != nil {
		return err
	}
	lock.session, err = concurrency.NewSession(lock.client,
		concurrency.WithLease(response.ID),
		concurrency.WithContext(lock.ctx))
	if err != nil {
		return err
	}
	lock.mutex = concurrency.NewMutex(lock.session, lock.key)
	if err = lock.mutex.TryLock(lock.ctx); err != nil {
		return err
	}
	return nil
}

func (lock *EtcdLock) UnLock() error {
	lock.cancel()
	_ = lock.session.Close()
	return lock.mutex.Unlock(context.TODO())
}

func (lock *EtcdLock) GetLockKey() string {
	return lock.key
}
