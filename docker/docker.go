package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

var rd *redis.Client

func main() {
	InitRedis()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "hello world")
	})
	http.HandleFunc("/set", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		key := request.Form.Get("key")
		val := request.Form.Get("val")
		fmt.Println("set >>>>>>  ", "key", key, "val", val)
		if key == "" {
			_, _ = writer.Write([]byte("参数错误"))
			return
		}
		if err := rd.Set(key, val, time.Second*60).Err(); err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		_, _ = writer.Write([]byte("操作成功"))
	})
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		key := request.Form.Get("key")
		fmt.Println("get >>>>>>  ", "key", key)
		if key == "" {
			_, _ = writer.Write([]byte("参数错误"))
			return
		}
		info, err := rd.Get(key).Result()
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		_, _ = writer.Write([]byte(info))
	})
	fmt.Println("服务启动成功 监听端口 9999")
	er := http.ListenAndServe("0.0.0.0:9999", nil)
	if er != nil {
		fmt.Println("ListenAndServe: ", er)
	}
}

func InitRedis() {
	rd = redis.NewClient(
		&redis.Options{
			Addr:         "redis:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "123456",
			PoolSize:     100,
		},
	)
	err := rd.Ping().Err()
	if nil != err {
		fmt.Println(err)
		return
	}
}
