package network

import (
	"net"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type TcpClient struct {
	Tag  string
	Conn net.Conn
	rw   *bufio.ReadWriter
	len  int32
}

func NewTcpClint(conn net.Conn) *TcpClient {
	return &TcpClient{Conn: conn, rw: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), len: 4}
}

func (c *TcpClient) Read() {
	defer c.Conn.Close()
	for {
		// 读取消息的长度
		lengthByte, err := c.rw.Peek(int(c.len))
		fmt.Println(lengthByte)
		if err == io.EOF {
			fmt.Println("用户退出")
			return
		}
		if len(lengthByte) == 0 || err != nil {
			fmt.Println(err)
			continue
		}
		lengthBuff := bytes.NewBuffer(lengthByte)
		var length int32
		err = binary.Read(lengthBuff, binary.LittleEndian, &length)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(length)
		if int32(c.rw.Reader.Buffered()) < length+c.len {
			fmt.Println(c.rw.Reader.Buffered())
			continue
		}
		// 读取消息真正的内容
		pack := make([]byte, int(c.len+length))
		_, err = c.rw.Read(pack)
		if err != nil {
			continue
		}
		//获取到的数据
		fmt.Println(string(pack[4:]))
	}
}
