package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok1"))
	})
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("服务器启动成功端口:", 3001)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Panic(err)
	}
}


