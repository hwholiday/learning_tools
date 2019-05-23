package main


//nats-streaming-server --store file -dir /home/ghost/data/nats
import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)
func main() {
	var clusterId string = "test-cluster"
	var clientId string = "test-client"
	sc, err := stan.Connect(clusterId, clientId, stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		for  {
			time.Sleep(time.Second)
			sc.Publish("foo", []byte("nast test"))
		}
	}()
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	defer sub.Unsubscribe()
	defer sc.Close()
	signalChan := make(chan int)
	<-signalChan

}