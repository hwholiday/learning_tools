package discovery

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/api/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"learning_tools/etcd/register"
	"log"
	"sync"
	"time"
)

type NodeArray struct {
	Node []register.Options `json:"node"`
}

type Discovery struct {
	etcdCli *clientv3.Client
	cc      resolver.ClientConn
	Node    sync.Map
	opts    *Options
}

func NewDiscovery(opt ...ClientOptions) resolver.Builder {
	s := &Discovery{
		opts: newOptions(opt...),
	}
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		panic(err)
	}
	s.etcdCli = etcdCli
	return s
}

// Build 当调用`grpc.Dial()`时执行
func (d *Discovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	d.cc = cc
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := d.etcdCli.Get(ctx, d.opts.SrvName, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, v := range res.Kvs {
		if err = d.AddNode(v.Key, v.Value); err != nil {
			log.Println(err)
			continue
		}
	}
	log.Printf("no %s service found , waiting for the service to join \n", d.opts.SrvName)
	go func(dd *Discovery) {
		dd.watcher()
	}(d)
	return d, err
}

func (d *Discovery) AddNode(key, val []byte) error {
	var data = new(register.Options)
	err := json.Unmarshal(val, data)
	if err != nil {
		return err
	}
	addr := resolver.Address{Addr: data.Node.Address}
	addr = SetNodeInfo(addr, data)
	d.Node.Store(string(key), addr)
	return d.cc.UpdateState(resolver.State{Addresses: d.GetAddress()})
}

func (d *Discovery) DelNode(key []byte) error {
	keyStr := string(key)
	d.Node.Delete(keyStr)
	return d.cc.UpdateState(resolver.State{Addresses: d.GetAddress()})
}

func (d *Discovery) GetAddress() []resolver.Address {
	var addr []resolver.Address
	d.Node.Range(func(key, value interface{}) bool {
		addr = append(addr, value.(resolver.Address))
		return true
	})
	return addr
}

func (d *Discovery) Scheme() string {
	return "discovery"
}

//watcher 监听前缀
func (d *Discovery) watcher() {
	rch := d.etcdCli.Watch(context.Background(), d.opts.SrvName, clientv3.WithPrefix())
	for res := range rch {
		for _, ev := range res.Events {
			switch ev.Type {
			case mvccpb.PUT: //新增或修改
				if err := d.AddNode(ev.Kv.Key, ev.Kv.Value); err != nil {
					log.Println(err)
				}
			case mvccpb.DELETE: //删除
				if err := d.DelNode(ev.Kv.Key); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (s *Discovery) ResolveNow(rn resolver.ResolveNowOptions) {
	//log.Println("ResolveNow")
}

func (s *Discovery) Close() {
	s.etcdCli.Close()
}
