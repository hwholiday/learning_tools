package match

import (
	"testing"
	"time"
)

func Test_NewMatchPool(t *testing.T)  {
	macth:=NewMatchPool(10000)
	macth.Add(&Match{
		Uid:       1,
		Rating:    1599,
		StartTime: time.Now().UnixNano()/1e6,
	})
	time.Sleep(time.Second/10)
	macth.Add(&Match{
		Uid:       2,
		Rating:    1600,
		StartTime: time.Now().UnixNano()/1e6,
	})
	time.Sleep(time.Second/10)
	macth.Add(&Match{
		Uid:       3,
		Rating:    1599,
		StartTime: time.Now().UnixNano()/1e6,
	})
	time.Sleep(time.Second/10)
	macth.Add(&Match{
		Uid:       4,
		Rating:    1599,
		StartTime: time.Now().UnixNano()/1e6,
	})
	time.Sleep(time.Second/10)
	macth.Add(&Match{
		Uid:       5,
		Rating:    1666,
		StartTime: time.Now().UnixNano()/1e6,
	})
	select {}
}