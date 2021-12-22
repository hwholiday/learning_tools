package hconfig

import (
	"github.com/hwholiday/learning_tools/hconfig/apollo"
	"github.com/hwholiday/learning_tools/hconfig/etcd"
	"github.com/hwholiday/learning_tools/hconfig/kubernetes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestNewHConfig_ETCD(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	c, err := etcd.NewEtcdConfig(cli,
		etcd.WithRoot("/hconf"),
		etcd.WithPaths("app", "mysql"))
	if err != nil {
		t.Error(err)
		return
	}
	conf, err := NewHConfig(
		WithDataSource(c),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Load(); err != nil {
		t.Error(err)
		return
	}
	val, err := conf.Get("app")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("val %+v\n", val.String())
	if err = conf.Watch(func(path string, v HVal) {
		t.Logf("path %s val %+v\n", path, v.String())

	}); err != nil {
		t.Error(err)
		return
	}
	select {}
}

func TestNewHConfig_K8S(t *testing.T) {
	cli, err := kubernetes.NewK8sClientset(kubernetes.KubeConfigPath("/home/app/conf/kube_config/local_kube.yaml"))
	if err != nil {
		t.Error(err)
		return
	}
	c, err := kubernetes.NewKubernetesConfig(cli, kubernetes.WithNamespace("im"), kubernetes.WithPaths("im-test-conf", "im-test-conf2"))
	if err != nil {
		t.Error(err)
		return
	}
	conf, err := NewHConfig(
		WithDataSource(c),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Load(); err != nil {
		t.Error(err)
		return
	}
	val, err := conf.Get("im-test-conf")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("val %+v\n", val.String())
	if err = conf.Watch(func(path string, v HVal) {
		t.Logf("path %s val %+v\n", path, v.String())

	}); err != nil {
		t.Error(err)
		return
	}
	select {}
}

func TestNewHConfig_Apollo(t *testing.T) {
	c, err := apollo.NewApolloConfig(
		apollo.WithAppid("test"),
		apollo.WithNamespace("test.yaml"),
		apollo.WithAddr("http://127.0.0.1:32001"),
		apollo.WithCluster("dev"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	conf, err := NewHConfig(
		WithDataSource(c),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Load(); err != nil {
		t.Error(err)
		return
	}
	val, err := conf.Get("test.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("val %+v\n", val.String())
	if err = conf.Watch(func(path string, v HVal) {
		t.Logf("path %s val %+v\n", path, v.String())

	}); err != nil {
		t.Error(err)
		return
	}
	select {}
}
