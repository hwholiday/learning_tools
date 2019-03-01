package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

func main() {
	cfg, err := ini.Load("./ini/my.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.Section("").Key("app_model").Value())
	fmt.Println(cfg.Section("mysql").Key("name").Value())
	fmt.Println(cfg.Section("mysql").Key("pass").Value())
}
