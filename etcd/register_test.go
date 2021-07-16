package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
)

func TestNewServiceRegister(t *testing.T) {
	var a int64 = 119152584752627712
	float, err := strconv.ParseFloat(fmt.Sprintf("%d", a), 64)
	fmt.Println(float)
	fmt.Println(err)

	return
	cli := NewEtcd(&Conf{
		Addr:        []string{"127.0.0.1:2379"},
		DialTimeout: 5,
	})
	s, err := NewRegister(cli, "1", "1", 10)
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	fmt.Println(s)
	go func() {
		if s.ListenLeaseRespChan() {
			c <- syscall.SIGQUIT
		}
	}()
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出")
			return
		default:
			return
		}
	}
}
