package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	CheckErr(err)
	psc := redis.PubSubConn{c}
	psc.Subscribe("howie")
	go func() {
		for {
			time.Sleep(3 * time.Second)
			publish("127.0.0.1:6379")
		}
	}()
	for {
		switch v := psc.Receive().(type) {
		case redis.Subscription:
			fmt.Printf("1  %s: %s %d\n", v.Channel, v.Kind, v.Count)
			break
		case redis.Message: //单个订阅subscribe
			fmt.Printf("2  %s: message: %s\n", v.Channel, v.Data)
			break
		case error:
			fmt.Println(v)
			break
		}
	}

}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func publish(redisServerAddr string) {
	c, err := redis.Dial("tcp",redisServerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	c.Do("PUBLISH", "howie", "hello")
}
