package main

import (
	"net/http"
	"github.com/inconshreveable/go-update"
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"fmt"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	v1 := g.Group("/v1").Use(v1Use)
	v1.GET("/ping", Ping)
	g.Handle("GET","/ping",Ping)
	s := &http.Server{
		Addr:           ":8072",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}

func v1Use(g *gin.Context)  {
	fmt.Println("请求开始")
	g.Next()
	fmt.Println("请求结束")
}

func Ping(g *gin.Context) {
	fmt.Println("请求中")
	g.String(http.StatusOK, "pong")
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}
	return nil
}
