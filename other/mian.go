package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	log.Println("服务器启动成功端口:",8070)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Panic(err)
	}
}
