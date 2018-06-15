package main

import (
	"net/http"
	"fmt"
	"test/push/handler"
	"flag"
	_"net/http/pprof"
)

var (
	Addr      = "127.0.0.1:8182"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8182", "addr")
	flag.Parse()
	go handler.H.Run()
	http.HandleFunc("/v1/push", handler.PushHandler)
	http.HandleFunc("/v1/report", handler.ReportHandler)
	fmt.Printf("Chat Run :%s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		fmt.Printf("WebSocker:%s", err.Error())
	}
}
