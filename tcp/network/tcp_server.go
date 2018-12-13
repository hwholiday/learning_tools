package network

import (
	"net"
	"github.com/hwholiday/libs/logtool"
	"go.uber.org/zap"
	"os"
	"time"
	"io"
	"fmt"
)

func InitTcp() {
	addr, err := net.ResolveTCPAddr("tcp", "192.168.2.28:8111")
	if err != nil {
		logtool.Zap.Error("create addr", zap.Error(err))
		os.Exit(0)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		logtool.Zap.Error("listen tcp", zap.Error(err))
		os.Exit(0)
	}
	logtool.Zap.Info("listen tcp", zap.String("地址", addr.String()))
	go acceptTcp(listener)
}

func acceptTcp(listener *net.TCPListener) {
	for {
		var (
			conn *net.TCPConn
			err  error
		)
		if conn, err = listener.AcceptTCP(); err != nil {
			logtool.Zap.Info("listener.Accept err", zap.Any(listener.Addr().String(), err))
			return
		}
		if err = conn.SetKeepAlive(false); err != nil {
			logtool.Zap.Info("conn.SetKeepAlive err", zap.Error(err))
			return
		}
		if err = conn.SetReadBuffer(1024); err != nil {
			logtool.Zap.Info("conn.SetReadBuffer err", zap.Error(err))
			return
		}
		if err = conn.SetWriteBuffer(1024); err != nil {
			logtool.Zap.Info("conn.SetWriteBuffer err", zap.Error(err))
			return
		}
		go serveTCP(conn)
	}
}

func serveTCP(conn *net.TCPConn) {
	client := NewTcpClint(conn, 4, 4)
	logtool.Zap.Debug("链接上来的用户", zap.Any("地址", client.RemoteAddr().String()))
	go func() {
		for {
			tag, data, err := client.Read()
			if err != nil {
				if err == io.EOF {
					logtool.Zap.Debug("用户断开链接", zap.Any("地址", client.RemoteAddr().String()))
				}
				client.conn.Close()
				return
			}
			message := make(chan byte, 1)
			go HeartBeating(client, message, 10)
			logtool.Zap.Info(fmt.Sprintf("客户端 : %s 传入类型", client.RemoteAddr().String()), zap.String(fmt.Sprintf("类型 : %d", tag), fmt.Sprintf("数据 : %s", string(data))))
			//做自己的业务逻辑
		}
	}()
}

func HeartBeating(client *TcpClient, bytes chan byte, timeout int) {
	select {
	case _ = <-bytes:
		client.conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(10 * time.Second):
		logtool.Zap.Debug("用户超时未发数据", zap.Any("地址", client.RemoteAddr().String()))
		client.conn.Close()
		break
	}
}
