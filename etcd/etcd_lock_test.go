package etcd

import (
	"fmt"
	"testing"
	"time"
)

func TestEtcdLock_TryLock(t *testing.T) {
	cli := NewEtcd(&Conf{
		Addr:        []string{"172.12.17.165:2379"},
		DialTimeout: 5,
	})
	lock := NewEtcdLock(cli, "/get/post", 5)

	if err := lock.TryLock(); err != nil {
		fmt.Println("1 err")
		return
	}
	fmt.Println("1 success")
	go fmt.Println(lock.TryLock())
	go fmt.Println(lock.TryLock())
	go fmt.Println(lock.TryLock())
	go fmt.Println(lock.TryLock())
	go fmt.Println(lock.TryLock())
	time.Sleep(5 * time.Second)
	fmt.Println(1, lock.UnLock())
	go func() {
		time.Sleep(6 * time.Second)
		if err := lock.TryLock(); err != nil {
			fmt.Println("2 err")
			return
		}
		fmt.Println("2 success")
	}()
	go func() {
		time.Sleep(6 * time.Second)
		if err := lock.TryLock(); err != nil {
			fmt.Println("3 err")
			return
		}
		fmt.Println("3 success")
	}()
	select {}
}
