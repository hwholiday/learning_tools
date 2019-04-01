package main

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	dis:=NewDisPatcher(2,4)
	dis.Run()
	for {
		dis.JobQueue <- Goods{Data: []byte(fmt.Sprint(time.Now().UnixNano()))}
	}
}
