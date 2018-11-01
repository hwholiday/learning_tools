package main

import (
	"testing"
	"net"
	"log"
	"fmt"
	"encoding/binary"
	"bytes"
	"time"
)

func Test(t *testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	for {
		data,_:=Encode("ping")
		time.Sleep(time.Second*4)
		_, err := conn.Write(data)
		fmt.Println(err)
	}
}
func Encode(message string) ([]byte, error) {
	// 读取消息的长度
	var length  = int32(len(message))
	var pkg  = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}