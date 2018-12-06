package main

import (
	"testing"
	"net"
	"log"
	"encoding/binary"
	"bytes"
	"time"
	"fmt"
	"bufio"
)

func Test(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.2.28:8111")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	go func() {
		for {
			data, _ := Encode("1")
			time.Sleep(time.Second * 4)
			_, err := conn.Write(data)
			fmt.Println(err)
		}
	}()

	reader := bufio.NewReader(conn)
	for {
		data, err := Read(reader)
		if err != nil {
			return
		}
		fmt.Println(string(data))
	}

}
func Encode(message string) ([]byte, error) {
	// 读取消息的长度
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Read(c *bufio.Reader) ([]byte, error) {
	lengthByte, err := c.Peek(4)
	if err != nil {
		return nil, err
	}
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}
	if int32(c.Buffered()) < length+4 {
		return nil, err
	}
	pack := make([]byte, int(4+length))
	_, err = c.Read(pack)
	if err != nil {
		return nil, err
	}
	return pack[4:], nil
}
