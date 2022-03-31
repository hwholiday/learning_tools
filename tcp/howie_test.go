package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func Test(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.2.28:8111")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	go func() {
		/*for {*/
		data, err := Encode("2")
		if err == nil {
			time.Sleep(time.Second * 4)
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
			}
		}

		/*}*/
	}()

	reader := bufio.NewReader(conn)
	for {
		tag, data, err := Read(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(tag)
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
	// 写入消息类型 最大为 0xFFFFFFF
	err = binary.Write(pkg, binary.BigEndian, int32(0x1))
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

func Read(c *bufio.Reader) (int32, []byte, error) {
	var headLen int32 = 4
	var tagLen int32 = 4
	lengthByte, err := c.Peek(int(headLen + tagLen))
	if err != nil {
		return 0, nil, err
	}
	var length int32
	lengthBuff := bytes.NewBuffer(lengthByte[:headLen])
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return 0, nil, err
	}
	var tag int32
	tagBuff := bytes.NewBuffer(lengthByte[headLen:])
	err = binary.Read(tagBuff, binary.BigEndian, &tag)
	if err != nil {
		return 0, nil, err
	}
	if int32(c.Buffered()) < length+headLen+tagLen {
		return 0, nil, err
	}
	pack := make([]byte, int(headLen+length+tagLen))
	_, err = c.Read(pack)
	if err != nil {
		return 0, nil, err
	}
	return tag, pack[headLen+tagLen:], nil
}
