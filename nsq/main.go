package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := nsq.NewConfig()
	//最大允许向两台NSQD服务器接受消息，默认是1，要特别注意
	config.MaxInFlight = 2
	c, err := nsq.NewConsumer("howie", "hwholiday", config)
	if err != nil {
		log.Println(1)
		log.Println(err)
		return
	}
	c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println(string(msg.Body))
		return nil
	}))
	if err = c.ConnectToNSQLookupd("0.0.0.0:4161"); err != nil {
		log.Fatalln(2)
		log.Fatalln(err)
		return
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
