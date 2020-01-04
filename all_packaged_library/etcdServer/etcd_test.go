package etcdServer

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

var etcdClientV3 *clientv3.Client

func init() {
	var err error
	etcdClientV3, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379/"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
}

var prefix = "/registry/server/"
var keyTtl int64 = 5

func TestV1(t *testing.T) {
	//创建节点
	for i := 1; i <= 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(keyTtl)*time.Second)
		grant, err := etcdClientV3.Grant(context.Background(), keyTtl)
		if err != nil {
			t.Log(err)
		}
		prefix += fmt.Sprint(i)
		_, err = etcdClientV3.Put(ctx, prefix, fmt.Sprintf("node_%d", i), clientv3.WithLease(grant.ID))
		if err != nil {
			t.Log(err)
		}
		keepAliveCh, err := etcdClientV3.KeepAlive(context.TODO(), grant.ID)
		if err != nil {
			t.Log(err)
		}
		if i == 3 {
			go func() {
				time.Sleep(time.Second * 20)
				_, err := etcdClientV3.Revoke(context.TODO(), grant.ID)
				if err != nil {
					t.Log(err)
				}
			}()
		}
		go func() {
			for range keepAliveCh {
			}
		}()
		fmt.Println("注册新节点 >>> prefix >>> ", prefix, "node >>>> ", fmt.Sprintf("node_%d", i))
		cancel()
		time.Sleep(time.Second * 10)
	}
	select {}
}
func TestV2(t *testing.T) {
	res, err := etcdClientV3.Get(context.TODO(), prefix, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		panic(err)
	}
	for _, kv := range res.Kvs {
		fmt.Println("key=", string(kv.Key), "|", "value=", string(kv.Value))
	}
	ch := etcdClientV3.Watch(context.TODO(), prefix, clientv3.WithPrefix())
	for {
		select {
		case c := <-ch:
			for _, e := range c.Events {
				fmt.Println(e.Type.String(), "key=", string(e.Kv.Key), "|", "value=", string(e.Kv.Value))
			}
		}
	}
}
