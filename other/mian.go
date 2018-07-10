package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
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

func main1() {
	fmt.Printf("%d\n", 1)
	fmt.Printf("%s\n", "1")
	fmt.Printf("%t\n", true)
	fmt.Printf("%.1f\n", 111.1111)
}
