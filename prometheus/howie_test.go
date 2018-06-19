package main

import (
	"testing"
	"fmt"
	"net/http"
	"time"
)

func Test(t *testing.T) {
	for {
		time.Sleep(50 * time.Millisecond)
		send()
	}
	time.Sleep(24*time.Hour)
}
func send() {
	resp, err := http.Get("http://localhost:8080/howie")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}
