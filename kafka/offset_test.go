package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
)

func TestKafkaSyncProducer(t *testing.T) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	conf.Producer.Retry.Max = 1000
	conf.Version = sarama.V2_8_0_0
	producer, err := sarama.NewSyncProducer([]string{"172.12.12.165:9092"}, conf)
	if err != nil {
		t.Error(err)
		return
	}
	defer producer.Close()
	fmt.Println(producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test_topic11",
		Value: sarama.ByteEncoder("1"),
	}))
	select {}
}

func TestConsumer(t *testing.T) {
	conf := sarama.NewConfig()
	conf.Version = sarama.V2_8_0_0
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := sarama.NewConsumerGroup([]string{"172.12.12.165:9092"}, "test_group", conf)
	if err != nil {
		t.Error(err)
		return
	}
	defer consumer.Close()
	ctx := context.Background()
	go func() {
		for {
			j := &Job{}
			if err := consumer.Consume(ctx, []string{"test_topic11"}, j); err != nil {
				fmt.Println("err", err)
			}
		}
	}()
	select {}
}

type Job struct {
}

func (consumer *Job) Setup(data sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Job) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Job) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Job Message claimed: value = %s, timestamp = %v, topic = %s , partition = %d , offset = %d",
			string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}
	return nil
}
