package timewheel

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	tw := DefaultTimeWheel()
	err := tw.AddTask("123123", func(i interface{}) {
		fmt.Println("11111")
	}, time.Second*5)
	t.Log("err", err)
	//time.Sleep(time.Second * 3)
	//tw.Stop()
	time.Sleep(time.Hour)
}
