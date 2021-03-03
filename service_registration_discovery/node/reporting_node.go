package library

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type Node struct {
	etcd    *clientv3.Client
	key     string
	Ip      string
	TTl     int64
	leaseId clientv3.LeaseID
}

func InitNode(etcdAddr []string, key, id, ip string, ttl int64) (n *Node, err error) {
	config := clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	}
	n = new(Node)
	n.key = fmt.Sprintf("/%s/%s", key, id)
	n.Ip = ip
	n.TTl = ttl
	n.etcd, err = clientv3.New(config)
	if err != nil {
		return
	}
	return
}
func (n *Node) UpNode() error {
	lease := clientv3.NewLease(n.etcd)
	leaseId, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		return err
	}
	n.leaseId = leaseId.ID
	_, err = n.etcd.Put(context.TODO(), n.key, n.Ip, clientv3.WithLease(n.leaseId))
	if err != nil {
		return err
	}
	_, err = lease.KeepAlive(context.TODO(), n.leaseId)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) Close() error {
	if _, err := n.etcd.Delete(context.TODO(), n.key); err != nil {
		return err
	}
	return nil
}
