package apollo

import (
	"testing"
)

func TestNewEtcdConfig(t *testing.T) {
	c, err := NewApolloConfig(
		WithAppid("test"),
		WithNamespace("test.yaml"),
		WithAddr("http://127.0.0.1:32001"),
		WithCluster("dev"),
	)
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
