package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		remote, err := url.Parse("http://127.0.0.1:8881")
		if err != nil {
			panic(err)
		}
		fmt.Println("accept request")
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(writer, request)
	})
	fmt.Println("proxy server success :", 8880)
	_ = http.ListenAndServe(":8880", mu)
}
