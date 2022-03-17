package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	go fmt.Println(doFuncSuccess(ctx))
	go fmt.Println(doFuncFail(ctx))
	time.Sleep(time.Second * 6)
	fmt.Println("doFuncFail")
	go fmt.Println(doFuncFail(ctx))
	time.Sleep(time.Second * 2)
}

func doFuncSuccess(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return ctx.Err().Error()
	default:
	}
	time.Sleep(time.Second * 3)
	select {
	case <-ctx.Done():
		return ctx.Err().Error()
	default:
	}
	return "success"
}

func doFuncFail(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return ctx.Err().Error()
	default:
	}

	time.Sleep(time.Second * 5)

	select {
	case <-ctx.Done():
		return ctx.Err().Error()
	default:
	}
	return "fail"
}
