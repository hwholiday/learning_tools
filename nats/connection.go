package main

import (
	"log"

	"github.com/nats-io/nats.go"
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
	}

	conn, err := natsOpt.Connect()
	if nil != err {
		log.Fatalln("failed to connect to nats-server", err)
	}
	defer conn.Close()

	log.Println("success to connect to nats-server,server_name: ", conn.ConnectedServerName())

}
