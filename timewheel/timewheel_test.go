package timewheel

import (
	"fmt"
	"testing"
)

func TestAddTask(t *testing.T) {
	tw := NewTimeWheel()
	tw.AddFn("123123", func(i interface{}) {
		fmt.Println("11111")
	}, 5)
	select {}
}
