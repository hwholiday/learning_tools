package main

import (
	"testing"
	"net/http"
)

func TestLimit(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func() {
			for {
				_, err := http.Get("http://192.168.2.28:8099/limit_api")
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	select {}
}
