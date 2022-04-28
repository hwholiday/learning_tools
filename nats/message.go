package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	serverAddr = "nats://localhost:4222"
	user       = "a"
	password   = "a"
	subject    = "test"
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

	// subscribe
	sub, err := conn.Subscribe(subject, func(msg *nats.Msg) {
		log.Println("onReader| Message Received: ", string(msg.Data))
		msg.Respond([]byte("reader has received your message"))
	})
	if nil != err {
		log.Fatalln("failed to establish subscription ", err)
	}
	sub.AutoUnsubscribe(1)

	// publish message
	log.Println("onWriter| prepare to send message")
	reply, err := conn.Request(subject, []byte("message from writer to reader"), 15*time.Second)
	if nil != err {
		log.Fatalln("onWriter| failed to send message to reader ", err)
	}
	log.Println("onWriter| success to send message to reader")
	log.Println("onWriter| response: ", string(reply.Data))

}
