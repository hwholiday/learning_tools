package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"testing"
	"time"
)

func TestRateLimitV1(t *testing.T) {
	limit := ratelimit.New(100) //每秒100个请求
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := limit.Take()
		if i == 3 {
			time.Sleep(time.Millisecond * 16)
		}
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

func TestRateLimitV2(t *testing.T) {
	limit := ratelimit.New(100, ratelimit.WithoutSlack) //每秒100个请求 WithoutSlack来取消松弛量
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := limit.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
