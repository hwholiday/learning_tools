package network

import (
	"net"
)

const (
	IP       = "ip"       // 链接的IP地址
	Identity = "identity" // 链接的UUID
)

type Conn interface {
	Conn() net.Conn
	Read() (byt []byte, err error)
	Write(byt []byte) error
	Cache() *Cache
	Close()
}
