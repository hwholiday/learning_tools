package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"sync"
)

type Node struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Address string `json:"address"`
}
type NodeInfo struct {
	Node      Node              `json:"node"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints map[string]string `json:"endpoints"`
}

type NodeArray struct {
	Node []NodeInfo `json:"node"`
}

type Discovery struct {
	etcdCli *clientv3.Client
	cc      resolver.ClientConn
	Node    sync.Map
	opts    *Options
}

type attributesEmpty struct{}

const scheme = "grpclb"

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
	fmt.Println("target", target)
	d.cc = cc
	res, err := d.etcdCli.Get(context.Background(), d.opts.SrvName, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, v := range res.Kvs {
		if err = d.AddNode(v.Key, v.Value); err != nil {
			return nil, err
		}
	}
	go d.watcher()
	return nil, err
}

func (d *Discovery) AddNode(key, val []byte) error {
	var data = new(NodeInfo)
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
	return scheme
}

//watcher 监听前缀
func (d *Discovery) watcher() {
	rch := d.etcdCli.Watch(context.Background(), d.opts.SrvName, clientv3.WithPrefix())
	for res := range rch {
		for _, ev := range res.Events {
			switch ev.Type {
			case mvccpb.PUT: //新增或修改
				d.AddNode(ev.Kv.Key, ev.Kv.Value)
			case mvccpb.DELETE: //删除
				d.DelNode(ev.Kv.Key)
			}
		}
	}
}

func (s *Discovery) Close() error {
	return s.etcdCli.Close()
}
