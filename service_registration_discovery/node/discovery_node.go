package library

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"strings"
	"time"
)

type DiscoveryNode struct {
	etcd  *clientv3.Client
	nodes map[string]string
	key   string
}

func InitDiscoveryNode(etcdAddr []string, key string) (n *DiscoveryNode, err error) {
	config := clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	}
	n = new(DiscoveryNode)
	n.key = key
	n.nodes = make(map[string]string)
	n.etcd, err = clientv3.New(config)
	if err != nil {
		return
	}
	n.watch()
	return
}

func (d *DiscoveryNode) GetNodes() map[string]string {
	return d.nodes
}

func (d *DiscoveryNode) watch() error {
	kv := clientv3.NewKV(d.etcd)
	rangeResp, err := kv.Get(context.TODO(), "/"+d.key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
		return err
	}
	for _, kv := range rangeResp.Kvs {
		d.nodes[d.getId(string(kv.Key))] = string(kv.Value)
	}
	go func() {
		curRevision := rangeResp.Header.Revision + 1
		watcher := clientv3.NewWatcher(d.etcd)
		watchChan := watcher.Watch(context.TODO(), d.key, clientv3.WithPrefix(), clientv3.WithRev(curRevision))
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.PUT:
					d.nodes[d.getId(string(event.Kv.Key))] = string(event.Kv.Value)
				case mvccpb.DELETE:
					delete(d.nodes, d.getId(string(event.Kv.Key)))
				}
			}
		}
	}()
	return nil
}

func (d *DiscoveryNode) getId(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(key, "/", ""), d.key, "")
}
