package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"testing"
	"time"
)

var mqProducer rocketmq.Producer

func TestMain(m *testing.M) {
	var err error
	fmt.Println("注册生产者")
	// 注册生产者
	if mqProducer, err = rocketmq.NewProducer(
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"172.13.3.160:9876"})),
		producer.WithRetry(3),
	); err != nil {
		panic(err)
	}
	mqProducer.Start()
	os.Exit(m.Run())
}

func TestMqOneWayProducer(t *testing.T) {
	err := mqProducer.SendOneWay(context.Background(), &primitive.Message{
		Topic: "msg",
		Body:  []byte("test"),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestMqSyncProducer(t *testing.T) {
	res, err := mqProducer.SendSync(context.Background(), &primitive.Message{
		Topic: "msg",
		Body:  []byte("Hello"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}

func TestMqASyncProducer(t *testing.T) {
	err := mqProducer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		fmt.Println("回调完", time.Now().UnixNano()/1e6)
		fmt.Println("result", result)
		fmt.Println("err", err)
	}, &primitive.Message{
		Topic: "msg",
		Body:  []byte("Hello SendAsync"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("执行完", time.Now().UnixNano()/1e6)
	select {}
}
func TestMqDelaySyncProducer(t *testing.T) {
	a := &primitive.Message{
		Topic: "msg",
		Body:  []byte("Hello TestMqDelaySyncProducer"),
	}
	a.WithTag("group")
	a.WithKeys([]string{"123"})
	a.WithDelayTimeLevel(3)
	res, err := mqProducer.SendSync(context.Background(), a)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}

func TestMqSubscribe(t *testing.T) {
	//注册消费者
	mqPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test_consumer"),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{"172.13.3.160:9876"})),
	)
	if err != nil {
		panic(err)
	}
	msgHandler := func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("接受到消息: %v /n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil

	}
	err = mqPushConsumer.Subscribe("msg", consumer.MessageSelector{}, msgHandler)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("启动消费者")
	mqPushConsumer.Start()
	select {}
}
