package main

import (
	"net/http"
	"log"
	"fmt"
)

func main1() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	log.Println("服务器启动成功端口:",8070)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Panic(err)
	}
}

func main(){
	fmt.Printf("%d\n",1)
	fmt.Printf("%s\n","1")
	fmt.Printf("%t\n",true)
	fmt.Printf("%.1f\n",111.1111)
}
