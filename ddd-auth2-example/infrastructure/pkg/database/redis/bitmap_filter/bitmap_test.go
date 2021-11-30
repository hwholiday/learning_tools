package bitmap_filter

import (
	"github.com/go-redis/redis"
	"testing"
)

func TestNewBitMapFileTer(t *testing.T) {
	r := redis.NewClient(&redis.Options{
		Addr: "172.12.12.165:6379",
	})
	b := NewBitMapFileTer(r, "gs_bitmap")
	t.Log(b.Exist("0x287dc8295e10878be9e20cb72e28c0e89bfe73b7"))
}
