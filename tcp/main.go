package main

import (
	"net"
	"test/tcp/network"
	"fmt"
)

func main() {
	lister, err := net.Listen("tcp", "127.0.0.1:8888")
	CheckErr(err)
	defer lister.Close()
	for {
		conn, err := lister.Accept()
		CheckErr(err)
		fmt.Println("用户接入")
		client := network.NewTcpClint(conn)
		go client.Read()
	}
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
