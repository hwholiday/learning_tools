package kubernetes

import "testing"

func TestNewKubernetesConfig(t *testing.T) {
	cli, err := NewK8sClientset(KubeConfigPath("/home/app/conf/kube_config/local_kube.yaml"))
	if err != nil {
		t.Error(err)
		return
	}
	c, err := NewKubernetesConfig(cli, WithNamespace("im"), WithPaths("im-test-conf", "im-test-conf2"))
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
