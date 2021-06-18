package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/groupcache/singleflight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var errPlaceholder = errors.New("placeholder")
var errNotFind = errors.New("not find")

type cache struct {
	rds *redis.Client
	g   singleflight.Group
}

/*create table tests
(
	id int auto_increment,
	user_name varchar(20) null,
	pwd varchar(30) null,
	create_time bigint null,
	update_time bigint null,
	constraint test_pk
	primary key (id)
);*/
//INSERT INTO test.tests (user_name, pwd, create_time, update_time) VALUES ('123123', '123123123', 123123123, 123123123);

type Test struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	UserName   string `json:"user_name"`
	Pwd        string `json:"pwd"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr: "172.12.12.165:6379",
	})
	sql := "root:+eN(2dFc5qu.@tcp(172.12.12.165:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"
	db, err := gorm.Open(mysql.Open(sql), nil)
	if err != nil {
		panic(err)
	}
	var c = &cache{
		rds: rds,
	}
	var test Test
	if err = c.Take(&test, "user:info:2", func(v interface{}) error {
		var t Test
		if err = db.Model(&Test{}).Where("id = ?", 2).Last(&t).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errNotFind
			} else {
				return err
			}
		}
		*v.(*Test) = t
		return nil
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test", test)
}

func (c *cache) Take(v interface{}, key string, query func(v interface{}) error) error {
	return c.Take2(v, key, query, func(v interface{}) error {
		//保存DB数据到缓存
		return c.Set(key, v)
	})
}

func (c *cache) Take2(v interface{}, key string, query func(v interface{}) error, cacheVal func(v interface{}) error) error {
	val, err := c.g.Do(key, func() (interface{}, error) {
		//从缓存里面读取
		if err := c.Get(key, v); err != nil {
			if err == errPlaceholder {
				return nil, errNotFind
			} else if err != errNotFind {
				return nil, err
			}
			//err == errNotFind
			//从DB里面读取
			if err = query(v); err == errNotFind {
				fmt.Println("query", err)
				if err = c.Set(key, "*"); err != nil {
					log.Println("set redis key  : ", key, " err :", err.Error())
				}
				return nil, errNotFind
			} else if err != nil {
				return nil, err
			}
			//调用保存缓存方法
			if err = cacheVal(v); err != nil {
				log.Println("cacheVal redis key  : ", key, " err :", err.Error())

			}
		}
		return json.Marshal(v)
	})
	if err != nil {
		return err
	}
	//等待到了 singleflight 的返回值并返回
	return json.Unmarshal(val.([]byte), v)
}

func (c *cache) Get(key string, v interface{}) error {
	data, err := c.rds.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return errNotFind
	}
	if err != nil {
		return err
	}
	if data == "\"*\"" {
		return errPlaceholder
	}
	if err = json.Unmarshal([]byte(data), v); err == nil {
		return nil
	}
	if err = c.rds.Del(key).Err(); err != nil {
		log.Println("del redis key  : ", key, " err :", err.Error())
	}
	return errNotFind
}

func (c *cache) Set(key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.rds.Set(key, data, time.Second*1160).Err()
}
