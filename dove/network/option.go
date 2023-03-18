package network

import (
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
	"net"
	"time"
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
	heartbeatInterval time.Duration
	autoHeartbeat     bool
}

func WithConn(conn net.Conn) Option {
	return func(o *options) {
		o.conn = conn
	}
}

func WithAutoHeartbeat(auto bool) Option {
	return func(o *options) {
		o.autoHeartbeat = auto
	}
}

func WithReadChanLen(len int) Option {
	return func(o *options) {
		o.readChanLen = len
	}
}

func WithWiterChanLen(len int) Option {
	return func(o *options) {
		o.witerChanLen = len
	}
}

func WithReadBufferSize(size int) Option {
	return func(o *options) {
		o.readBufferSize = size
	}
}

func WithWiterBufferSize(size int) Option {
	return func(o *options) {
		o.witerBufferSize = size
	}
}

func WithHeartbeatInterval(t time.Duration) Option {
	return func(o *options) {
		o.heartbeatInterval = t
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
		length:            8,
		heartbeatInterval: time.Second * 30,
		useBigEndian:      true,
		autoHeartbeat:     true,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.conn == nil {
		return nil, errors.New("conn is nil")
	}
	if o.useBigEndian {
		o.endian = binary.BigEndian
	} else {
		o.endian = binary.LittleEndian
	}
	return o, nil
}
