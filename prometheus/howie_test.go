package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	for {
		time.Sleep(50 * time.Millisecond)
		send()
	}
	time.Sleep(24 * time.Hour)
}
func send() {
	resp, err := http.Get("http://192.168.2.28:8888/howie")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}
