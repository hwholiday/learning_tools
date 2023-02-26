package timewheel

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	tw := DefaultTimeWheel()
	err := tw.AddTask("task-1", func(key string) {
		fmt.Println("task run :", key, " > ", time.Now().Format(time.DateTime))
	}, 3*time.Second, 6)
	t.Log("err", err)
	//time.Sleep(time.Second * 3)
	//tw.Stop()
	//tw.RemoveTask("task-1")
	time.Sleep(time.Hour)
}
