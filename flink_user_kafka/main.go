package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//异步
	AsyncProducer, err := sarama.NewAsyncProducer([]string{"172.12.17.161:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer AsyncProducer.AsyncClose()
	go func() {
		for {
			select {
			case succ := <-AsyncProducer.Successes():
				fmt.Println("AsyncProducer.Successes()", succ.Topic, succ.Offset, succ.Timestamp, succ.Partition)
			case err := <-AsyncProducer.Errors():
				fmt.Println("AsyncProducer.Errors()", err.Error())
			}
		}
	}()
	var i = 1
	for {
		Message := &sarama.ProducerMessage{
			Topic:     "user_login",
			Timestamp: time.Now(),
		}
		var data []byte
		date := time.Now().Unix()
		if i == 1 {
			data = []byte(fmt.Sprintf("%s:%d:%d", "hw", date, 1))
			i = 2
		} else {
			data = []byte(fmt.Sprintf("%s:%d:%d", "hr", date, 1))
			i = 1
		}

		Message.Value = sarama.ByteEncoder(data)
		AsyncProducer.Input() <- Message
		time.Sleep(time.Second * 1)
	}
}

type Login struct {
	Username  string `json:"username"`
	LoginTime int64  `json:"loginTime"`
}
