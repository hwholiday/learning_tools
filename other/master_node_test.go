package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func Test_NodeA(t *testing.T) {
	Node("127.0.0.1:1111")
}
func Test_NodeB(t *testing.T) {
	Node("127.0.0.1:2222")
}
func Test_NodeC(t *testing.T) {
	Node("127.0.0.1:3333")
}
func Test_NodeD(t *testing.T) {
	Node("127.0.0.1:4444")
}
func Test_NodeE(t *testing.T) {
	Node("127.0.0.1:5555")
}
func Node(ip string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		fmt.Println(ip, "New", err)
		return
	}
	var master bool
	var ID clientv3.LeaseID
	lease := clientv3.NewLease(cli)
	for {
		if !master {
			txn := clientv3.NewKV(cli).Txn(context.TODO())
			grantRes, err := lease.Grant(context.TODO(), 10) //创建一个10秒的租约
			if err != nil {
				fmt.Println(ip, "Grant", err)
				master = false
				continue
			}
			ID = grantRes.ID
			txn.If(clientv3.Compare(clientv3.CreateRevision("/id/master"), "=", 0)).
				Then(clientv3.OpPut("/id/master", ip, clientv3.WithLease(grantRes.ID))).
				Else()
			txnResp, err := txn.Commit()
			if err != nil {
				fmt.Println(ip, "Commit", err)
				master = false
				continue
			}
			if txnResp.Succeeded {
				fmt.Println(ip, "主节点")
				master = true
			} else {
				fmt.Println(ip, "从节点")
				master = false
			}
		}
		_, err = lease.KeepAliveOnce(context.TODO(), ID)
		if err != nil {
			fmt.Println(ip, "Put", err)
			continue
		}
		time.Sleep(9 * time.Second)
	}
	select {}
}
