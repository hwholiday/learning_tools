package main

import (
	"net/http"
	"time"
	"log"
	"test/jwt/controller"
)

func main() {
	mux := http.NewServeMux()
	token := &controller.TokenController{}
	mux.HandleFunc("/create_token", token.CreateToken)
	mux.HandleFunc("/test_token", token.TestToken)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("服务器启动成功:8080")
	err := s.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
