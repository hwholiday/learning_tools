package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"gopkg.in/yaml.v2"
)

const (
	serverAddr = "nats://localhost:4222"
	user       = "a"
	password   = "a"
)

/*
 * nats-server -js -sd . -D --user a --pass a
 * docker run --rm -ti -p 4222:4222 nats:2.8.1-alpine -js -sd . -D --user a --pass a
 */

func main() {
	natsOpt := &nats.Options{
		Url:            serverAddr,
		Name:           "test-client",
		AllowReconnect: true,
		User:           user,
		Password:       password,
	}

	conn, err := natsOpt.Connect()
	if nil != err {
		log.Fatalln("failed to connect to nats-server", err)
	}
	defer conn.Close()

	log.Println("success to connect to nats-server,server_name: ", conn.ConnectedServerName())

	// Get jetstream context
	// js, err := conn.JetStream(nats.Domain("test"))
	js, err := conn.JetStream()
	if nil != err {
		log.Fatalln("failed to connect to jetstream ", err)
	}
	log.Println("success to get jetstream")

	// Create stream => any message sent to subjects filterd by stream would be persisted
	var (
		streamName      string = "test"
		filteredSubject string = "test.>" // capture test.1 test.2 test.1.1,etc...
	)
	streamCfg := &nats.StreamConfig{
		Name:        streamName,
		Description: "this is a test subject",
		Subjects:    []string{filteredSubject},
		Retention:   nats.WorkQueuePolicy,
		Discard:     nats.DiscardNew,
		Storage:     nats.FileStorage,
		Replicas:    1,
	}

	stream, err := js.AddStream(streamCfg)
	if nil != err {
		log.Fatalln("js failed to add stream ", err)
	}
	log.Println("js success to add stream ", stream.Config.Name)
	streamCfgData, _ := yaml.Marshal(stream.Config)
	log.Printf("create time: %v\nconfig:\n%s", stream.Created, string(streamCfgData))

	// Send Message to stream
	sendSubject := "test.1"
	resp, err := js.Publish(sendSubject, []byte("this is a test message"))
	if nil != err {
		log.Fatalln("js failed to publish message ", err)
	}
	log.Println("success to publish message: ", resp)

	// Get stream info
	streamInfo, err := js.StreamInfo(streamName)
	if nil != err {
		log.Fatalln("js failed to  get stream info ", err)
	}
	log.Printf("streamName: %s\ncurrent messages: %d", streamInfo.Config.Name, streamInfo.State.Msgs)

	// Create Consumer in Push Mode
	var (
		consumerName            string = "test1"
		consumerFilteredSubject string = "test.1"
		deliverSubject          string = "receive.1"
	)
	consumerCfg := &nats.ConsumerConfig{
		Durable:        consumerName,
		Description:    "this is a test consumer for subject 'test.1' in stream test",
		DeliverSubject: deliverSubject,
		DeliverPolicy:  nats.DeliverAllPolicy,
		AckPolicy:      nats.AckExplicitPolicy,
		FilterSubject:  consumerFilteredSubject,
		ReplayPolicy:   nats.ReplayInstantPolicy,
		MaxAckPending:  20,
	}

	consumer, err := js.AddConsumer(streamName, consumerCfg)
	if nil != err {
		log.Fatalln("js failed to add consumer ", err)
	}
	log.Println("js success to add consumer")
	consumerCfgData, _ := yaml.Marshal(consumer.Config)
	log.Printf("created at %v\nconfig:\n%v", consumer.Created, string(consumerCfgData))
	log.Println("current messages: ", consumer.NumPending)

	// Get messages from consumer
	var wg sync.WaitGroup
	wg.Add(1)
	sub, err := js.Subscribe(consumerFilteredSubject, func(msg *nats.Msg) {
		log.Println("onReader|  message received: ", string(msg.Data))
		msg.Ack()
		wg.Done()
	}, nats.Bind(streamName, consumerName))
	if nil != err {
		log.Fatalln("failed to establish subscription ", err)
	}
	log.Println("success to establish subscription")
	sub.AutoUnsubscribe(1)

	wg.Wait()

	// Get consumer info
	consumerInfo, err := js.ConsumerInfo(streamName, consumerName)
	if nil != err {
		log.Fatalln("failed to get consumer info ", err)
	}
	log.Printf("consumer name: %s\ncurrent messages: %d", consumerInfo.Name, consumerInfo.NumPending)

}
