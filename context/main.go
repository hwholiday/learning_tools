package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var siganChannel = make(chan os.Signal, 1)

func main() {
	//ContextWithCancel()
	//ContextWithTimeout()
	ContextWithDeadline()
}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Value("howie"))
				return
			}
		}
	}()
	fmt.Println("开始")
	time.AfterFunc(10*time.Second, func() {
		ctx = context.WithValue(ctx, "howie", "10秒后调用cancel()")
		cancel()
		fmt.Println("结束")
	})
	Exit()

}

func eeContextWithTimeout() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(time.Since(start))
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
				return
			}
		}
	}()
	Exit()

}

func ContextWithDeadline() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	start := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(time.Since(start))
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
				return
			}
		}
	}()
	Exit()

}

func Exit() {
	signal.Notify(siganChannel, os.Kill, os.Interrupt)
	<-siganChannel
}
