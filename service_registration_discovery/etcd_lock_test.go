package service_registration_discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/v3/clientv3"
	"go.etcd.io/etcd/v3/clientv3/concurrency"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	config := clientv3.Config{
		Endpoints:   []string{"172.13.3.160:2379"},
		DialTimeout: 5 * time.Second,
	}
	etcd, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	response, err := etcd.Grant(timeout, 5)
	if err != nil {
		panic(err)
	}
	session, err := concurrency.NewSession(etcd, concurrency.WithLease(response.ID))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	for {
		time.Sleep(time.Second)
		mu := concurrency.NewMutex(session, "/user/order")
		err := mu.Lock(timeout)
		if err != nil {
			fmt.Println("获取锁成功")
		} else {
			fmt.Println("获取锁失败")
		}
	}

	//defer mu.Unlock(timeout)
}
