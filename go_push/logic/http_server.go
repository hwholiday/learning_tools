package logic

import (
	"fmt"
	"net/http"
	"time"
)

func InitHttpServer() {
	//curl  http://0.0.0.0:9999/push/all -X POST -d "val=test push&tag=3"
	//curl  http://0.0.0.0:9999/push/room -X POST -d "val=test push&tag=3&id=1"
	//curl http://0.0.0.0:9999/room/join -X POST -d "id=1&wsId=3cc97117-aa4d-55bb-86db-a1d77e51283b"
	//curl http://0.0.0.0:9999/room/leave -X POST -d "id=1&wsId=3cc97117-aa4d-55bb-86db-a1d77e51283b"
	mux := http.NewServeMux()
	mux.HandleFunc("/push/all", HttpPushAll)
	mux.HandleFunc("/push/room", HttpPushRoom)
	mux.HandleFunc("/room/join", HttpRoomJoin)
	mux.HandleFunc("/room/leave", HttpRoomLeave)
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
