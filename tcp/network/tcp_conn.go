package network

import (
	"net"
	"bufio"
	"bytes"
	"encoding/binary"
)

type TcpClient struct {
	tag  string
	conn net.Conn
	rw   *bufio.ReadWriter
}

func NewTcpClint(conn net.Conn) *TcpClient {
	return &TcpClient{conn: conn, rw: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))}
}

func (c *TcpClient) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TcpClient) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TcpClient) Close() error {
	return c.conn.Close()
}

func (c *TcpClient) Write(message []byte) ([]byte, error) {
	// 读取消息的长度
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err = binary.Write(pkg, binary.LittleEndian, message)
	if err != nil {
		return nil, err
	}
	_, err = c.rw.Write(pkg.Bytes())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *TcpClient) Read() ([]byte, error) {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	lengthByte, err := c.rw.Peek(4)
	if err != nil {
		return nil, err
	}
	//创建 Buffer缓冲器
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	// 通过Read接口可以将buf中得内容填充到data参数表示的数据结构中
	err = binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	// Buffered 返回缓存中未读取的数据的长度
	if int32(c.rw.Reader.Buffered()) < length+4 {
		return nil, err
	}
	// 读取消息真正的内容
	pack := make([]byte, int(4+length))
	// Read 从 b 中读出数据到 p 中，返回读出的字节数和遇到的错误。
	// 如果缓存不为空，则只能读出缓存中的数据，不会从底层 io.Reader
	// 中提取数据，如果缓存为空，则：
	// 1、len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读
	// 出到 p 中。
	// 2、len(p) < 缓存大小，则先将数据从底层 io.Reader 中读取到缓存
	// 中，再从缓存读取到 p 中。
	_, err = c.rw.Read(pack)
	if err != nil {
		return nil, err
	}
	return pack[4:], nil
}
