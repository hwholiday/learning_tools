package match

import (
	"testing"
	"time"
)

func Test_NewMatchPool(t *testing.T) {
	macth := NewMatchPool(100000)
	macth.Add(&Match{
		Uid:       1,
		Rating:    1599,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	time.Sleep(time.Second)
	macth.Add(&Match{
		Uid:       2,
		Rating:    1600,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	time.Sleep(time.Second)
	macth.Add(&Match{
		Uid:       3,
		Rating:    1599,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	time.Sleep(time.Second)
	macth.Add(&Match{
		Uid:       4,
		Rating:    1599,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	time.Sleep(time.Second)
	macth.Add(&Match{
		Uid:       5,
		Rating:    1666,
		StartTime: time.Now().UnixNano() / 1e6,
	})

	select {}
}

func TestMatchPool_AddV1(t *testing.T) {
	var middle = 100
	for searchRankUp, searchRankDown := middle, middle; searchRankUp <= 120 && searchRankDown >= 80 ; searchRankUp, searchRankDown = searchRankUp+1, searchRankDown-1 {
		t.Log("searchRankUp: ", searchRankUp)
		t.Log("searchRankDown: ", searchRankDown)
		if searchRankDown != searchRankUp && searchRankDown > 0 {
			t.Log("searchRankDown: searchRankDownsearchRankDownsearchRankDown ", searchRankDown)
		}
	}

}
