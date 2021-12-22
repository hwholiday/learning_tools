package etcd

import (
	"testing"
)
import clientv3 "go.etcd.io/etcd/client/v3"

func TestNewEtcdConfig(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	c, err := NewEtcdConfig(cli, WithRoot("/hconf"), WithPaths("app", "mysql"))
	if err != nil {
		t.Error(err)
		return
	}
	loadData, err := c.Load()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range loadData {
		t.Logf("Load key %s val %s \n", v.Key, string(v.Val))
	}
	w, err := c.Watch()
	defer w.Close()
	if err != nil {
		t.Error(err)
		return
	}
	go func() {
		for {
			kvs, err := w.Change()
			if err != nil {
				return
			}
			for _, v := range kvs {
				t.Logf("Change key %s val %s \n", v.Key, string(v.Val))
			}
		}
	}()
	select {}
}
