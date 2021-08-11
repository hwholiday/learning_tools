package bitmap_filter

import (
	"github.com/go-redis/redis"
	"testing"
)

func Test_hashKey(t *testing.T) {
	d := "0x0a0f824b9e1e951560d9fa1fd4b89403d092d046"
	t.Log(hashKey(d))
}

func TestNewBitMapFileTer(t *testing.T) {
	r := redis.NewClient(&redis.Options{
		Addr: "172.12.12.165:6379",
	})
	b := NewBitMapFileTer(r, "gs_bitmap")
	t.Log(b.Add("0x0a0f824b9e1e951560d9fa1fd4b89403d092d046"))
	t.Log(b.Exist("0x0a0f824b9e1e951560d9fa1fd4b89403d092d0461"))
}
