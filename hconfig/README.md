### [hconfig 插拔式配置读取工具可动态加载](https://github.com/hwholiday/learning_tools/tree/master/hconfig)

- 支持 etcd
- 支持 kubernetes
- 支持 apollo

#### hconfig  配置不同的源

```base
//etcd
cli, err := clientv3.New(clientv3.Config{
	Endpoints: []string{"127.0.0.1:2379"},})
	
c, err := etcd.NewEtcdConfig(cli,
	etcd.WithRoot("/hconf"),
	etcd.WithPaths("app", "mysql"))

//kubernetes
cli, err := kubernetes.NewK8sClientset(
     kubernetes.KubeConfigPath("/home/app/conf/kube_config/local_kube.yaml"))
     
c, err := kubernetes.NewKubernetesConfig(cli, 
	kubernetes.WithNamespace("im"),
	kubernetes.WithPaths("im-test-conf", "im-test-conf2"))
	
//apollo
c, err := apollo.NewApolloConfig(
    apollo.WithAppid("test"),
    apollo.WithNamespace("test.yaml"),
    apollo.WithAddr("http://127.0.0.1:32001"),
    apollo.WithCluster("dev"),
    )
```


#### hconfig 使用

```base
conf, err := NewHConfig(
	WithDataSource(c),//c 不同的源
)

// 加载配置
conf.Load() 

//读取配置
val, err := conf.Get("test.yaml")
t.Logf("val %+v\n", val.String())

//监听配置变化
conf.Watch(func(path string, v HVal) {
	t.Logf("path %s val %+v\n", path, v.String())
})
```