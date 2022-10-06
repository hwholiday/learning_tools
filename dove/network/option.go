package network

import (
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
	"net"
)

type Option func(*options)

type options struct {
	id                string
	conn              net.Conn
	useBigEndian      bool
	endian            binary.ByteOrder
	length            int
	readBufferSize    int
	witerBufferSize   int
	witerChanLen      int
	readChanLen       int
	HeartbeatInterval int //s
	AutoHeartbeat     bool
}

func WithConn(conn net.Conn) Option {
	return func(o *options) {
		o.conn = conn
	}
}

func WithLength(length int) Option {
	return func(o *options) {
		o.length = length
	}
}

func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func newOptions(opts ...Option) (*options, error) {
	o := &options{
		id:                uuid.New().String(),
		witerBufferSize:   4096,
		readBufferSize:    4096,
		witerChanLen:      1,
		readChanLen:       1,
		HeartbeatInterval: 30,
		useBigEndian:      true,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.conn == nil {
		return nil, errors.New("conn is nil")
	}
	if o.length == 0 {
		return nil, errors.New("length is 0")
	}
	if o.useBigEndian {
		o.endian = binary.BigEndian
	} else {
		o.endian = binary.LittleEndian
	}
	return o, nil
}
