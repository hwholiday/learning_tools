package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Config struct {
	Addr        string `yaml:"addr"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Active      uint64 `yaml:"active"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

func NewMongo(conf *Config) *mongo.Client {
	opt := options.Client().ApplyURI(conf.Addr)
	if len(conf.User) != 0 { // 部分连接不需要帐号密码
		opt.Auth = &options.Credential{
			Username: conf.User,
			Password: conf.Password,
		}
	}
	opt.SetLocalThreshold(3 * time.Second)                                //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(time.Duration(conf.IdleTimeout) * time.Second) //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(conf.Active)                                       //使用最大的连接数
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	return client
}
