package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/toml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/etcd"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	t := toml.NewEncoder()
	sou := etcd.NewSource(
		etcd.WithAddress("192.168.1.86:2379"),
		etcd.WithPrefix("/conf/"),
		etcd.StripPrefix(true),
		source.WithEncoder(t),
	)
	if err = conf.Load(sou); err != nil {
		panic(err)
	}
	fmt.Println("读取到配置", string(conf.Get("test", "server").Bytes()))
	wath, err := conf.Watch("test", "server")
	if err != nil {
		panic(err)
	}
	defer wath.Stop()
	for {
		val, err := wath.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(val.Bytes()))
	}
}
