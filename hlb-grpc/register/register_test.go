package register

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNewServiceRegister(t *testing.T) {
	s, err := NewRegister(
		SetName("hwholiday.srv.msg"),
		SetAddress("127.0.0.1:123123"),
		SetVersion("v1"),
		SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"172.12.12.165:2379"},
			DialTimeout: time.Second * 5,
		}),
	)
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	go func() {
		if s.ListenKeepAliveChan() {
			c <- syscall.SIGQUIT
		}
	}()
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for a := range c {
		switch a {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出")
			_ = s.Close()
			return
		default:
			return
		}
	}
}
