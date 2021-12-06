package hevent

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var ch = make(chan HEvent)
	var ch2 = make(chan HEvent)
	HEventSrv().Sub("hw", ch)
	HEventSrv().Sub("hw", ch2)
	//go GetHEventSrv().Push("topic1", "Hi topic 1")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			HEventSrv().Push("hw", "Hi topic 1")
		}
	}()
	for {
		select {
		case c := <-ch:
			fmt.Println("ch", c)
		case c := <-ch2:
			fmt.Println("ch2", c)
		}
	}
}
