package main

import (
	"context"
	"github.com/go-redis/redis"
)

type Client struct {
	redis.Cmdable
}

type Config struct {
	PoolSize int
	Addr     []string
	Pwd      string
}

func NewRedis(o Config) (client *Client, err error) {
	var redisCli redis.Cmdable
	if len(o.Addr) > 1 {
		redisCli = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:    o.Addr,
				PoolSize: o.PoolSize,
				Password: o.Pwd,
			},
		)
	} else {
		redisCli = redis.NewClient(
			&redis.Options{
				Addr:     o.Addr[0],
				Password: o.Pwd,
				PoolSize: o.PoolSize,
				DB:       0,
			},
		)
	}
	err = redisCli.Ping(context.TODO()).Err()
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
		return redisCli.Process(context.TODO(), cmd)
	case *redis.Client:
		return redisCli.Process(context.TODO(), cmd)
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
