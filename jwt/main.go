package main

import (
	"net/http"
	"time"
	"log"
	"test/jwt/controller"
)

func main() {
	//Http.request.body.read（）方法
	http.Request{}.Body.Read()

	mux := http.NewServeMux()
	token := &controller.TokenController{}
	mux.HandleFunc("/create_token", token.CreateToken)
	mux.HandleFunc("/create_rsa_token", token.CreateTokenByRsa)
	mux.HandleFunc("/test_token", token.TestToken)
	mux.HandleFunc("/test_rsa_token", token.TestRsaToken)
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
