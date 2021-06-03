package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Conf struct {
	Addr        []string
	DialTimeout int
}

func NewEtcd(c *Conf) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Addr,
		DialTimeout: time.Duration(c.DialTimeout) * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}
