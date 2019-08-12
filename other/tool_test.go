package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
)

//获取本地的IP
func Test_GetIp(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(addrs)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

//设置配置文件路径
func Test_Path(t *testing.T) {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	fmt.Println(appConfigPath)
}

func Test_Defer(t *testing.T)  {
	 defer fmt.Println(1)
	 A()
	defer fmt.Println(3)
	return
}
func A()  {
	defer func() {
		fmt.Println(2)
	}()
}
