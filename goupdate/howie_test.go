package main

import (
	"testing"
	"os/exec"
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/inconshreveable/go-update"
	"time"
	"log"
	"os"
	"io/ioutil"
	"io"
)

func Test(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	v1 := g.Group("/v1").Use(v1Use)
	v1.GET("/ping", Ping)
	g.Handle("GET", "/version", Version)
	s := &http.Server{
		Addr:           ":8072",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("服务启动端口 :8072")
	if err := s.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}

func v1Use(g *gin.Context) {
	fmt.Println("请求开始")
	g.Next()
	fmt.Println("请求结束")
}

func Ping(g *gin.Context) {
	fmt.Println("请求中")
	g.String(http.StatusOK, "pong")
	return
}

func Version(g *gin.Context) {
	url := g.Query("url")
	fmt.Println(11)
	if url == "" {
		fmt.Println(1)
		g.String(http.StatusOK, "url为空")
		return
	}
	dir, err := os.Getwd()
	fmt.Println(22)
	if err != nil {
		fmt.Println(2)
		g.String(http.StatusOK, err.Error())
		return
	}
	file, err := ioutil.ReadFile(dir + "/pid.txt")
	fmt.Println(33)
	if err != nil {
		fmt.Println(3)
		g.String(http.StatusOK, err.Error())
		return
	}
	pid := string(file)
	fmt.Println(44)
	err, _ = execShell("kill -9 " + pid)
	if err != nil {
		fmt.Println(4)
		g.String(http.StatusOK, err.Error())
		return
	}
	resp, err := http.Get(url)
	fmt.Println(55)
	if err != nil {
		fmt.Println(5)
		g.String(http.StatusOK, err.Error())
		return
	}
	defer resp.Body.Close()
	name := fmt.Sprintf("main_%v", time.Now().Unix())
	f, err := os.Create(fmt.Sprintf("%v/%s", dir, name))
	fmt.Println(66)
	if err != nil {
		fmt.Println(6)
		g.String(http.StatusOK, err.Error())
		return
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	fmt.Println(77)
	if err != nil {
		fmt.Println(7)
		g.String(http.StatusOK, err.Error())
		return
	}
	fmt.Println(88)
	err, _ = execShell("chmod +x " + name)
	if err != nil {
		fmt.Println(8)
		g.String(http.StatusOK, err.Error())
		return
	}
	fmt.Println(99)
	go execServer("./" + name)
	g.String(http.StatusOK, "ok")
	return
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	fmt.Println("开始下载")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("下载完毕")
	err = update.Apply(resp.Body, update.Options{
		TargetPath: "/home/howie/go/src/test/goupdate/",
	})
	if err != nil {
		return err
	}
	return nil
}
func execShell(s string) (err error, data string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	data = out.String()
	return
}

func execServer(str string) {
	cmd := exec.Command(str)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}else{
		cmd.Wait()
	}
	fmt.Println(str)
}