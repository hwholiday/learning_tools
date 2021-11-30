package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Client struct {
	redis.Cmdable
}

type Config struct {
	PoolSize     int           `yaml:"poolSize"`
	Addr         []string      `yaml:"addr"`
	Pwd          string        `yaml:"pwd"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

func NewRedis(o *Config) (client *Client, err error) {
	var redisCli redis.Cmdable
	if len(o.Addr) > 1 {
		redisCli = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:        o.Addr,
				PoolSize:     o.PoolSize,
				DialTimeout:  o.DialTimeout,
				ReadTimeout:  o.ReadTimeout,
				WriteTimeout: o.WriteTimeout,
				Password:     o.Pwd,
			},
		)
	} else {
		redisCli = redis.NewClient(
			&redis.Options{
				Addr:         o.Addr[0],
				DialTimeout:  o.DialTimeout,
				ReadTimeout:  o.ReadTimeout,
				WriteTimeout: o.WriteTimeout,
				Password:     o.Pwd,
				PoolSize:     o.PoolSize,
				DB:           0,
			},
		)
	}
	err = redisCli.Ping().Err()
	if nil != err {
		panic(err)
	}

	client = new(Client)
	client.Cmdable = redisCli
	return client, nil
}

func (c *Client) Process(cmd redis.Cmder) error {
	switch redisCli := c.Cmdable.(type) {
	case *redis.ClusterClient:
		return redisCli.Process(cmd)
	case *redis.Client:
		return redisCli.Process(cmd)
	default:
		return nil
	}
}

func (c *Client) Close() error {
	switch redisCli := c.Cmdable.(type) {
	case *redis.ClusterClient:
		return redisCli.Close()
	case *redis.Client:
		return redisCli.Close()
	default:
		return nil
	}
}
