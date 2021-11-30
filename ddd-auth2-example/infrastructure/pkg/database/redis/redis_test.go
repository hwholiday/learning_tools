package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	fmt.Println("1", time.Now().Format("2006-01-02 15:04:05"))
	cli, err := NewRedis(&Config{
		PoolSize: 5,
		Addr: []string{
			"172.12.12.165:6379",
		},
		DialTimeout:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
	})
	if err != nil {
		panic(err)
	}
	cli.ZAdd("test_z", redis.Z{
		Score:  1,
		Member: "A",
	})
	cli.ZAdd("test_z", redis.Z{
		Score:  2,
		Member: "B",
	})
	cli.ZAdd("test_z", redis.Z{
		Score:  3,
		Member: "C",
	})
	cli.ZAdd("test_z", redis.Z{
		Score:  4,
		Member: "D",
	})
	cli.ZAdd("test_z", redis.Z{
		Score:  5,
		Member: "F",
	})
	cacheMsgKey := make([]string, 0)
	offset := "+inf"
	err = cli.ZRevRangeByScore("test_z", redis.ZRangeBy{
		Min:    "-inf",
		Max:    offset,
		Offset: 0,
		Count:  0,
	}).ScanSlice(&cacheMsgKey)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cacheMsgKey)
}
