package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

func InitTcp() {
	addr, err := net.ResolveTCPAddr("tcp", "192.168.2.28:8111")
	if err != nil {
		os.Exit(0)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		os.Exit(0)
	}
	go acceptTcp(listener)
}

func acceptTcp(listener *net.TCPListener) {
	for {
		var (
			conn *net.TCPConn
			err  error
		)
		if conn, err = listener.AcceptTCP(); err != nil {
			return
		}
		if err = conn.SetKeepAlive(false); err != nil {
			return
		}
		if err = conn.SetReadBuffer(1024); err != nil {
			return
		}
		if err = conn.SetWriteBuffer(1024); err != nil {
			return
		}
		go serveTCP(conn)
	}
}

func serveTCP(conn *net.TCPConn) {
	client := NewTcpClint(conn, 4, 4)
	client.conn.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))
	go func() {
		for {
			tag, data, err := client.Read()
			if err != nil {
				client.Close()
				return
			}
			fmt.Println("tag", tag)
			fmt.Println("data", data)
			client.conn.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))
			//做自己的处理
		}
	}()
}
