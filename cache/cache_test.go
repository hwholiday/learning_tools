package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetSize(t *testing.T) {
	result, err := GetSize("100KB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 102400, result, "获取内存大小")
}

func TestCache_Set(t *testing.T) {
	cache := NewCache(time.Second * 10)
	cache.SetMaxMemory("100KB")
	cache.Set("1", "2", time.Second*5)
	val, ok := cache.Get("1")
	if !ok {
		t.Fatal("err")
	}
	cache.Get("1")
	cache.Get("1")
	cache.Get("1")
	if fmt.Sprint(val) != "2" {
		t.Fatal("set err")
	}
	time.Sleep(time.Second * 6)
	val, ok = cache.Get("1")
	if ok {
		t.Fatal("err")
	}
}
