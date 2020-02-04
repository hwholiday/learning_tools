package main

import (
	"fmt"
	"net/http"
)

func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("test request")
		_, _ = writer.Write([]byte("test success"))
	})
	fmt.Println("server success :", 8881)
	_ = http.ListenAndServe(":8881", mu)
}
