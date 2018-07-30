package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"flag"
	"time"
	"github.com/cihub/seelog"
	"os"
	"fmt"
	"test/gin/router"
	"github.com/go-xorm/xorm"
	"test/gin/model"
)

var addr = flag.String("addr", "127.0.0.1:8081", "server addr")
var seelogConfig = flag.String("log", "conf/seelog.xml", "seelog config")
var mysqlPath = flag.String("mysql", "conf/mysql.json", "mysql config")
var Xengine *xorm.Engine
func init() {
	logger, err := seelog.LoggerFromConfigAsFile(*seelogConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	seelog.ReplaceLogger(logger)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	router.SetRouters(g)
	model.InitDb(*mysqlPath)
	s := &http.Server{
		Handler:        g,
		Addr:           *addr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	seelog.Info("server run :", *addr)
	err := s.ListenAndServe()
	if err != nil {
		seelog.Error(err.Error())
		seelog.Flush()
		os.Exit(0)
	}
}
