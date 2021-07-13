package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli := NewEtcd(&Conf{
		Addr:        []string{"172.12.12.165:2379"},
		DialTimeout: 5,
	})
	kv := clientv3.NewKV(cli)
	// 先GET到当前的值，并监听后续变化
	getResp, err := kv.Get(context.TODO(), "/conf", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		return
	}

	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}

	// 获得当前revision
	watchStartRevision := getResp.Header.Revision + 1
	// 创建一个watcher
	watcher := clientv3.NewWatcher(cli)
	fmt.Println("从该版本向后监听:", watchStartRevision)
	watchRespChan := watcher.Watch(context.TODO(), "/conf", clientv3.WithPrefix(), clientv3.WithRev(watchStartRevision))
	// 处理kv变化事件
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("key", string(event.Kv.Key), "修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("key", string(event.Kv.Key), "删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}
}
