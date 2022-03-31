package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

type My struct {
	AppModel string `ini:"app_model"`
	Type     int    `ini:"type"`
	Sql      Mysql  `ini:"mysql"`
}
type Mysql struct {
	Name string `ini:"name"`
	Pass string `ini:"pass"`
}

func main() {
	cfg, err := ini.Load("./ini/my.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg.Section("").Key("app_model").Value())
	fmt.Println(cfg.Section("mysql").Key("name").Value())
	fmt.Println(cfg.Section("mysql").Key("pass").Value())
	p := new(My)
	err = cfg.MapTo(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}
