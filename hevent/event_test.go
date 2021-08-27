package hevent

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var ch = make(chan HEvent, 10)
	HEventSrv().Sub("hw", ch)
	//go GetHEventSrv().Push("topic1", "Hi topic 1")
	go func() {
		for {
			time.Sleep(1000 * time.Second)
			HEventSrv().Push("hw", "Hi topic 1")
		}
	}()
	for {
		fmt.Println(<-ch)
	}
}
