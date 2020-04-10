package match

import (
	"testing"
	"time"
)

func Test_NewMatchPool(t *testing.T) {
	macth := NewMatchPool(100000, 2)
	macth.Add(&Match{
		Uid:       1,
		Rating:    1661,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	time.Sleep(time.Second)
	macth.Add(&Match{
		Uid:       2,
		Rating:    1662,
		StartTime: time.Now().UnixNano() / 1e6,
	})
	macth.Run()
	time.Sleep(time.Hour)
}

func TestMatchPool_AddV1(t *testing.T) {
	var middle = 100
	for searchRankUp, searchRankDown := middle, middle; searchRankUp <= 120 && searchRankDown >= 80; searchRankUp, searchRankDown = searchRankUp+1, searchRankDown-1 {
		t.Log("searchRankUp: ", searchRankUp)
		t.Log("searchRankDown: ", searchRankDown)
		if searchRankDown != searchRankUp && searchRankDown > 0 {
			t.Log("searchRankDown: searchRankDownsearchRankDownsearchRankDown ", searchRankDown)
		}
	}

}
