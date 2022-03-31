package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"
)

func Test(t *testing.T) {
	i := 0
	for {
		i++
		if i == 100000 {
			fmt.Println("已经了解10000个链接")
			break
		}
		time.Sleep(time.Second)
		go func(id int) {
			u := url.URL{Scheme: "ws", Host: "127.0.0.1:8182", Path: "/v1/push", RawQuery: fmt.Sprintf("uid=%d&sign=2&time=%d", id, time.Now().Unix())}
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Printf("链接地址%s,失败原因%s", u.String(), err.Error())
				return
			}
			defer c.Close()
			/*go func() {
				for {
					time.Sleep(time.Second * 30)
					data, err := json.Marshal(handler.ClientsReport{id, 1, "hahh"})
					if err != nil {
						log.Println("ERR_ERR_ERR_ERR_ERR_ERR_write11111111:", err)
						continue
					}
					err = c.WriteMessage(websocket.TextMessage, data)
					if err != nil {
						log.Println("ERR_ERR_ERR_ERR_ERR_ERR_write:", err)
						continue
					}

				}
			}()*/
			for {
				_, _, err := c.ReadMessage()
				if err != nil {
					log.Println("ERR_ERR_ERR_ERR_ERR_ERR_read:", err)
					continue
				}
			}
		}(i)
	}
	log.Println("启动模拟测试开始")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	s := <-ch
	fmt.Println("信号：", s)
}
