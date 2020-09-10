package gcache

import (
	"testing"
)

func TestNewGCache(t *testing.T) {
	lru := NewLRU(0, nil)
	lru.Add("1", []byte("1"))
	val, ok := lru.Get("1")
	if !ok {
		t.Fatalf("lru not get")
	}
	if string(val) != "1" {
		t.Fatalf("lru  get val err")
	}
	t.Log("val : ", string(val))
	lru.Del("1")
	_, ok = lru.Get("1")
	if ok {
		t.Fatalf("lru del err")
	}
}
