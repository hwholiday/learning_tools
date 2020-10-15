package test_agent

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"testing"
	"time"
)

func Lock() {
	config := clientv3.Config{
		Endpoints:   []string{"172.13.3.160:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	//创建一个20秒过期的锁
	session, err := concurrency.NewSession(client, concurrency.WithTTL(20))
	if err != nil {
		panic(err)
	}
	timeout, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	m := concurrency.NewMutex(session, "test-mutex")
	//10秒内不能获取到锁就不再获取
	if err := m.Lock(timeout); err != nil {
		fmt.Println("could not wait on lock ", err)
	} else {
		fmt.Println("获取到锁")
	}
}

func TestConcurrency(t *testing.T) {
	go Lock()
	go Lock()
	go Lock()
	select {}
}
