package main

import (
	"net/http"
	"log"
)

func main1() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("服务器启动成功端口:", 8070)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Panic(err)
	}
}


