package main

import (
	"github.com/go-redis/redis"
	"testing"
)

func Test_singleflight_group(t *testing.T) {
	redis.NewZSliceCmd()
}
