package network

import (
	"net"
	"github.com/hwholiday/libs/logtool"
	"go.uber.org/zap"
	"os"
	"io"
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
	client := NewTcpClint(conn)
	logtool.Zap.Debug("链接上来的用户", zap.Any("地址", client.RemoteAddr().String()))
	go func() {
		for {
			data, err := client.Read()
			if err == io.EOF {
				logtool.Zap.Debug("用户断开链接", zap.Any("地址", client.RemoteAddr().String()))
				return
			}
			if err != nil {
				continue
			}
			SwitchController(data, client)
		}
	}()

}
