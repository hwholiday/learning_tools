package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

var (
	_               Conn = (*conn)(nil)
	AlreadyCloseErr      = errors.New("conn already close")
)

// conn length + data 模式
type conn struct {
	opts       *options
	readWriter *bufio.ReadWriter
	cache      *Cache
	stopChan   chan bool
	writerChan chan []byte
	readChan   chan []byte
	once       sync.Once
	tmp        any
}

func NewConn(opt ...Option) (Conn, error) {
	var (
		err error
		c   = getConn()
	)
	c.opts, err = newOptions(opt...)
	if err != nil {
		return nil, err
	}
	c.readWriter = bufio.NewReadWriter(
		bufio.NewReaderSize(c.opts.conn, c.opts.readBufferSize),
		bufio.NewWriterSize(c.opts.conn, c.opts.witerBufferSize))
	c.stopChan = make(chan bool)
	c.cache = NewCache()
	c.writerChan = make(chan []byte, c.opts.readChanLen)
	c.readChan = make(chan []byte, c.opts.witerChanLen)
	c.cache.Save(IP, c.opts.conn.RemoteAddr().String())
	c.cache.Save(Identity, c.opts.id)
	go c.readChannel()
	go c.witerChannel()
	return c, nil
}

func (c *conn) Cache() *Cache {
	return c.cache
}

func (c *conn) Close() {
	c.once.Do(func() {
		c.stopChan <- true
		close(c.stopChan)
		close(c.readChan)
		close(c.writerChan)
		_ = c.opts.conn.Close()
		putConn(c)
	})
}

func (c *conn) Read() (byt []byte, err error) {
	select {
	case byt = <-c.readChan:
		return byt, nil
	case <-c.stopChan:
		return nil, AlreadyCloseErr
	}
}

func (c *conn) Conn() net.Conn {
	return c.opts.conn
}
func (c *conn) ResetConnDeadline() error {
	return c.opts.conn.SetDeadline(time.Now().Add(time.Duration(c.opts.HeartbeatInterval) * time.Second))
}

func (c *conn) Write(byt []byte) error {
	select {
	case c.writerChan <- byt:
	case <-c.stopChan:
		return AlreadyCloseErr
	}
	return nil
}

func (c *conn) readChannel() {
	for {
		byt, err := c.read()
		if err != nil {
			log.Printf("[Dove] readChannel Close conn id : %s , err: %s \n", c.opts.id, err.Error())
			c.Close()
			return
		}
		_ = c.ResetConnDeadline()
		select {
		case c.readChan <- byt:
		case <-c.stopChan:
			return
		}
	}
}

func (c *conn) witerChannel() {
	for {
		select {
		case byt := <-c.writerChan:
			if err := c.witer(byt); err != nil {
				log.Printf("[Dove] witerChannel id : %s , err: %s \n", c.opts.id, err.Error())
			}
		case <-c.stopChan:
			break
		}
	}
}

func (c *conn) read() ([]byte, error) {
	lengthByte, err := c.readWriter.Reader.Peek(c.opts.length)
	if err != nil {
		return nil, err
	}
	var length int
	if err = binary.Read(bytes.NewReader(lengthByte), c.opts.endian, &length); err != nil {
		return nil, err
	}

	if c.readWriter.Reader.Buffered() < int(c.opts.length+length) {
		return nil, errors.New("the corresponding data cannot be read")
	}
	pack := make([]byte, int(c.opts.length+length))

	if _, err = c.readWriter.Reader.Read(pack); err != nil {
		return nil, err
	}
	return pack[c.opts.length:], err
}

func (c *conn) witer(byt []byte) error {
	var (
		length = len(byt)
	)
	if err := binary.Write(c.readWriter.Writer, c.opts.endian, length); err != nil {
		return err
	}
	if err := binary.Write(c.readWriter.Writer, c.opts.endian, byt); err != nil {
		return err
	}
	return c.readWriter.Writer.Flush()
}
