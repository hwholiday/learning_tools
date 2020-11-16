package main

import "net/http"

func main() {
	http.HandleFunc("/k8s", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world k8s"))
	})
	_ = http.ListenAndServe(":8080", nil)
}
