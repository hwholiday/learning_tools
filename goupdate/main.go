package main

import (
	"net/http"
	"log"
	"os"
	"strconv"
	"fmt"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	fmt.Println(os.Getpid())
	dir, _ := os.Getwd()
	file, _ := os.Create(dir + "/pid.txt")
	fmt.Println(dir + "/pid.txt")
	file.WriteString(strconv.Itoa(os.Getpid()))
	file.Close()
	log.Println("服务器启动成功端口:", 8071)
	err := http.ListenAndServe(":8071", nil)
	if err != nil {
		log.Panic(err)
	}
}
