package registrySelector

import (
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

/*func TestV2(t *testing.T) {
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
}*/

func TestNewRegistry(t *testing.T) {
	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"http://localhost:2379/"},
			DialTimeout: 5 * time.Second},
	}
	r, err := NewRegistry(op)
	if err != nil {
		t.Error(err)
		return
	}
	err = r.RegistryNode(Node{Addr: "127.0.0.1:1001", Id: "1"})
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 20)
	//r.UnRegistry()
	time.Sleep(time.Hour * 5)
}
