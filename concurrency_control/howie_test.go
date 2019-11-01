package main

import (
	"net/http"
	"testing"
)

func TestLimit(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func() {
			for {
				_, err := http.Get("http://192.168.2.28:8099/limit_api_v1")
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	select {}
}
