package main

import (
	"testing"
	"github.com/nsqio/go-nsq"
	"time"
	"log"
)

func TestPr(t *testing.T) {
	send("howie","0.0.0.0:4150")
	send("howie","0.0.0.0:4152")
}

func send(tag string, addr string) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Println(err)
		return
	}
	p.Publish(tag,[]byte(tag+":"+time.Now().String()))
}