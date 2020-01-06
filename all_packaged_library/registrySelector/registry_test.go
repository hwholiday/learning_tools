package registrySelector

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestNewRegistry(t *testing.T) {
	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"http://localhost:2379/"},
			DialTimeout: 5 * time.Second},
	}
	for i := 1; i <= 3; i++ {
		r, err := NewRegistry(op)
		if err != nil {
			t.Error(err)
			return
		}
		err = r.RegistryNode(PutNode{Addr: fmt.Sprintf("127.0.0.1:%d%d%d%d", i, i, i, i)})
		if err != nil {
			t.Error(err)
			return
		}
		if i == 3 {
			go func() {
				time.Sleep(time.Second * 20)
				r.UnRegistry()
			}()
		}

	}
	time.Sleep(time.Hour * 5)
}

func TestNewSelector(t *testing.T) {
	var op = SelectorOptions{
		name: "svc.info",
		config: clientv3.Config{
			Endpoints:   []string{"http://localhost:2379/"},
			DialTimeout: 5 * time.Second},
	}
	s, err := NewSelector(op)
	if err != nil {
		t.Error(err)
		return
	}
	for {
		val, err := s.Next()
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}
}