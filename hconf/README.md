### 基于etcd与viper的高可用配置中心

- 可使用远程与本地模式
- 本地有的配置远程没有会自动把本地配置传到远程（基于key）
- 远程有的配置本地没有也会写一份到本地(退出程序会把远程配置写一份到本地)
- 远程模式配置可以动态加载
- 如远程连接不上会使用本地配置启动作为兜底

```base
var conf = Conf{}
r, err := NewHConf(
	SetWatchRootName([]string{"/gs/conf"}),
)
if err != nil {
	t.Error(err)
	return
}
t.Log(r.ConfByKey("/gs/conf/net", &conf.Net))
t.Log(r.ConfByKey("/gs/conf/net2222", &conf.Net2))
t.Log(r.ConfByKey("/gs/conf/net3333", &conf.Net3))
if err := r.Run(); err != nil {
	t.Error(err)
	return
}
t.Log(conf)
t.Log(r.Close())
```
