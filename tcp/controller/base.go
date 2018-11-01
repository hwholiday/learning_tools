package controller

import (
	"fmt"
	"test/tcp/network"
	"io"
	"net"
)

func ServerRun() {
	lister, err := net.Listen("tcp", "192.168.2.28:8888")
	fmt.Println("服务启动成功：127.0.0.1:8888")
	CheckErr(err)
	defer lister.Close()
	for {
		conn, err := lister.Accept()
		CheckErr(err)
		fmt.Println("用户接入")
		client := network.NewTcpClint(conn)
		go func() {
			defer client.Close()
			for {
				data, err := client.Read()
				if err == io.EOF {
					fmt.Println("断开链接")
					return
				}
				if err != nil {
					continue
				}
				switchController(data, client)
			}
		}()
	}
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func switchController(data []byte, c *network.TcpClient) {
	fmt.Println("读到的数据: " + string(data))
	switch string(data) {
	case "ping":
		c.Write([]byte("pong"))
		fmt.Println("发出的数据: pong")
		break
	}
}
