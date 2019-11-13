package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

func Test_Producer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	AsyncProducer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	CheckErr(err)
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
	for {
		Message := &sarama.ProducerMessage{
			Topic:     "Test_Producer",
			Timestamp: time.Now(),
		}
		value := "消息测试 Test_Producer " + time.Now().Format("2006-01-02 15:04:05")
		Message.Value = sarama.ByteEncoder(value)
		AsyncProducer.Input() <- Message
		time.Sleep(time.Second * 10)
	}
}

func Test_Consumer(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	CheckErr(err)
	defer consumer.Close()
	client, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	CheckErr(err)
	defer client.Close()
	offset, err := sarama.NewOffsetManagerFromClient("test_group", client)
	CheckErr(err)
	defer offset.Close()
	offsetManager, err := offset.ManagePartition("Test_Producer", 0)
	CheckErr(err)
	defer offsetManager.Close()
	nextOffset, _ := offsetManager.NextOffset()
	fmt.Println(nextOffset)
	PartitionConsumer, err := consumer.ConsumePartition("Test_Producer", 0, nextOffset+1)
	CheckErr(err)
	defer PartitionConsumer.Close()
	for {
		select {
		case data := <-PartitionConsumer.Messages():
			fmt.Println("Value", string(data.Value))
			fmt.Println("Timestamp", data.Timestamp)
			fmt.Println("Offset", data.Offset)
			fmt.Println("Topic", data.Topic)
			offsetManager.MarkOffset(data.Offset, "") //设置偏移量
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
