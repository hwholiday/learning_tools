package gateway

import (
	"fmt"
	"net/http"
	"time"
)

func InitHttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/push/all", HttpPushAll)
	mux.HandleFunc("/push/room", HttpPushRoom)
	// HTTP服务
	server := http.Server{
		Addr:         "0.0.0.0:9999",
		ReadTimeout:  time.Duration(10) * time.Millisecond,
		WriteTimeout: time.Duration(10) * time.Millisecond,
		Handler:      mux,
	}
	fmt.Println("启动HTTP服务器成功 ：", 9999)
	_ = server.ListenAndServe()
}
