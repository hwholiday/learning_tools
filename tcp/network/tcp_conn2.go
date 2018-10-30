package network

import (
	"net"
	"encoding/binary"
	"io"
)

type TcpClient2 struct {
	Tag    string
	len    int32
}

func NewTcpClint2(conn *net.TCPConn) *TcpClient {
	return &TcpClient{Conn: conn, len: 4}
}
func (c *TcpClient) Read2() ([]byte, error) {
	var b [c.len]byte
	bufMsgLen := b[:c.len]
	// read len
	if _, err := io.ReadFull(c.Conn, bufMsgLen); err != nil {
		return nil, err
	}
	var msgLen uint32
	msgLen = uint32(binary.BigEndian.Uint16(bufMsgLen))
	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(c.Conn, msgData); err != nil {
		return nil, err
	}
	return msgData, nil
}
