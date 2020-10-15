package service_registration_discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

var etcd *clientv3.Client

func init() {
	var err error
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	etcd, err = clientv3.New(config)
	CheckErr(err)
}

//存值
func TestPut(t *testing.T) {
	_, err := etcd.Put(context.TODO(), "/info/a", "127.0.0.1:2222")
	CheckErr(err)
	_, err = etcd.Put(context.TODO(), "/info/b", "127.0.0.1:2222")
	CheckErr(err)
	t.Log("put success")
}

//取值
func TestGet(t *testing.T) {
	data, err := etcd.Get(context.TODO(), "/info/a") //取指定key
	CheckErr(err)
	t.Log(data)
	data, err = etcd.Get(context.TODO(), "/info/", clientv3.WithPrefix()) //取带有/info/前缀的key
	CheckErr(err)
	t.Log(data)
	data, err = etcd.Get(context.TODO(), "/server/file") //取带有/info/前缀的key
	CheckErr(err)
	t.Log(data)
}

//创建一个5秒的租约 实现服务注册
func TestPutWithGrant(t *testing.T) {
	lease := clientv3.NewLease(etcd)
	leaseId, err := lease.Grant(context.TODO(), 5)
	CheckErr(err)
	_, err = etcd.Put(context.TODO(), "/server/file", "8.8.8.8", clientv3.WithLease(leaseId.ID))
	CheckErr(err)
	t.Log("put success")
	data, err := lease.KeepAlive(context.TODO(), leaseId.ID) //自动续约
	t.Log("[lease.KeepAlive]   ", data)
	go WatchData()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		data, err := etcd.Get(context.TODO(), "/server/file") //取指定key
		CheckErr(err)
		/*if i == 4 { //要过期的时候续一次租约
			_, err = lease.KeepAliveOnce(context.TODO(), leaseId.ID)
			CheckErr(err)
		}*/
		if i == 7 { //删除这个key
			_, err = lease.Revoke(context.TODO(), leaseId.ID)
			CheckErr(err)
			_, err = etcd.Delete(context.TODO(), "/server/file")
			CheckErr(err)
		}
		t.Log("第 ", i, " 秒获取数据", data)

	}
}

//实现服务发现  发现所有服务前缀为/server/的服务
func WatchData() {
	kv := clientv3.NewKV(etcd)
	rangeResp, err := kv.Get(context.TODO(), "/server/", clientv3.WithPrefix())
	CheckErr(err)
	for _, kv := range rangeResp.Kvs {
		fmt.Println("Key >> ", string(kv.Key), "Value >> ", string(kv.Value))
	}
	// 从当前版本开始订阅
	curRevision := rangeResp.Header.Revision + 1
	watcher := clientv3.NewWatcher(etcd)
	watchChan := watcher.Watch(context.TODO(), "/server/", clientv3.WithPrefix(), clientv3.WithRev(curRevision))
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("PUT事件 Key >> ", string(event.Kv.Key), "Value >> ", event.Kv.Value)
			case mvccpb.DELETE:
				fmt.Println("DELETE事件 Key >> ", string(event.Kv.Key))
			}
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


