package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
	"time"
)

func Test_Queue(t *testing.T) {
	//同步
	SyncProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	CheckErr(err)
	defer SyncProducer.Close()
	SyncProducer.SendMessage(&sarama.ProducerMessage{})
}

func Test_Producer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//异步
	AsyncProducer, err := sarama.NewAsyncProducer([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
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
			Topic:     "ghost_topic",
			Timestamp: time.Now(),
		}
		value := "消息测试 ghost_topic " + time.Now().Format("2006-01-02 15:04:05")
		Message.Value = sarama.ByteEncoder(value)
		AsyncProducer.Input() <- Message
		time.Sleep(time.Second * 10)
	}
}
func Test_ConsumerGroup(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, "test_group", config)
	CheckErr(err)
	consumer := Consumer1{}
	for {
		if err := consumerGroup.Consume(context.Background(), []string{"ghost_topic"}, &consumer); err != nil {
			panic(err)
		}
	}
}

func Test_ConsumerGroup2(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, "test_group_1", config)
	CheckErr(err)
	consumer := Consumer{}
	for {
		if err := consumerGroup.Consume(context.Background(), []string{"ghost_topic"}, &consumer); err != nil {
			panic(err)
		}
	}
}

type Consumer struct {
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Consumer Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}

type Consumer1 struct {
}

func (consumer *Consumer1) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer1) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer1) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Consumer1 Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
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
