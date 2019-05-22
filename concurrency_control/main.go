package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var addr = flag.String("p", "192.168.2.28:8099", "port")

func main() {
	flag.Parse()
	//http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/limit_api", limitApi)
	s := &http.Server{
		Handler:        mux,
		Addr:           *addr,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

var limitChan = make(chan bool, 1) //每次执行一个请求
func limitApi(w http.ResponseWriter, r *http.Request) {
	select {
	case limitChan <- true:
	case <-time.After(2 * time.Second): //2秒不能写入服务器返回错误
		fmt.Println("服务器限流")
		w.Write([]byte("ERR"))
		return
	}
	//延迟释放limitChan 模拟超时
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	<-limitChan
	fmt.Println("服务器正常访问")
	w.Write([]byte("OK"))
	return
}
