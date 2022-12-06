#  LEARNING_TOOLS
## 博客: [hwholiday](https://www.hwholiday.com/posts/)
- 🤪 golang 基础，微服务，架构，web，k8s,DDD 各种工具任君取用
- 📫 contact me: WECHAT HW_loner
- 📫 contact me: QQ 3355168235





# go-kit 微服务实践，从入门到精通

### [go-kit 系列文章归档地址](https://github.com/hwholiday/learning_tools/tree/master/go-kit) (详细介绍)

1:  [v1 go-kit 微服务 基础使用 （HTTP）](https://www.hwholiday.com/2019/go_kit_v1/)  
2:  [v2 go-kit 微服务 添加日志（user/zap ,并为每个请求添加UUID）](https://www.hwholiday.com/2020/git_kit_v2/)   
3:  [v3 go-kit 微服务 身份认证 （JWT）](https://www.hwholiday.com/2020/go_kit_v3/)  
4:  [v4 go-kit 微服务 限流 （uber/ratelimit 和 golang/rate 实现）](https://www.hwholiday.com/2020/go_kit_v4/)  
5:  [v5 go-kit 微服务 使用GRPC（并为每个请求添加UUID）](https://www.hwholiday.com/2020/go_kit_v5/)   
6:  [v6 go-kit 微服务 服务注册与发现（etcd实现）](https://www.hwholiday.com/2020/go_kit_v6/)  
7:  [v7 go-kit 微服务 服务监控（prometheus 实现）](https://www.hwholiday.com/2020/go_kit_v7/)  
8:  [v8 go-kit 微服务 服务熔断（hystrix-go 实现）](https://www.hwholiday.com/2020/go_kit_v8/)  
9:  [v9 go-kit 微服务 服务链路追踪（jaeger 实现）(1)](https://www.hwholiday.com/2020/go_kit_v9/)  
10: [v10 go-kit 微服务 服务链路追踪（jaeger 实现）(2)](https://www.hwholiday.com/2020/go_kit_v10/)  
11: [v11 go-kit 微服务 日志分析管理 （ELK + Filebeat）](https://www.hwholiday.com/2020/go_kit_v12/)  

# gRPC负载均衡（自定义负载均衡策略--etcd 实现）

### [hlb-grpc](https://github.com/hwholiday/learning_tools/tree/master/hlb-grpc) (gRPC负载均衡（自定义负载均衡策略--etcd 实现)

##### 实现基于版本（version）的grpc负载均衡器，了解过程后可自己实现更多的负载均衡功能

### [详细介绍](https://www.hwholiday.com/2021/etcd_grpc/)

+ 注册中心
    - Etcd Lease 是一种检测客户端存活状况的机制。 群集授予具有生存时间的租约。 如果etcd 群集在给定的TTL 时间内未收到keepAlive，则租约到期。 为了将租约绑定到键值存储中，每个key 最多可以附加一个租约
+ 服务注册 (注册服务)
    - 定时把本地服务（APP）地址,版本等信息注册到服务器
+ 服务发现 (客户端发起服务解析请求（APP）)
    - 查询注册中心（APP）下有那些服务
    - 并向所有的服务建立HTTP2长链接
    - 通过Etcd watch 监听服务（APP），通过变化更新链接
+ 负载均衡 (客户端发起请求（APP）)
    - 负载均衡选择合适的服务（APP HTTP2长链接）
    - 发起调用

```
├── discovery
│   ├── customize_balancer.go
│   ├── discovery.go
│   └── options.go
├── example
│   ├── api
│   │   └── api.pb.go
│   ├── api.proto
│   ├── client_test.go
│   └── server.go
└── register
    ├── options.go
    ├── register.go
    └── register_test.go

```

# 仿微信 auth2 授权登陆 DDD（领域设计驱动）+ 六边形架构
### [OAuth 2.0-授权码模式（authorization code）仿微信设计（战略篇）](https://www.hwholiday.com/2022/auth2_strategy/)
### [OAuth 2.0-授权码模式（authorization code）仿微信设计（战术篇）](https://www.hwholiday.com/2022/auth2_tactics/)

### [AUTH2 代码地址](https://github.com/hwholiday/learning_tools/tree/master/ddd-auth2-example)

# Golang DDD 的项目分层结构（六边形架构）

### [DDD](https://github.com/hwholiday/learning_tools/tree/master/ddd-project-example) （DDD 项目分层结构）

```base
├── cmd 存放 main.go 等
├── adapter
│   ├── grpc
│   └── http
│   └── facade  引用其他微服务（接口防腐层）
├── application 
│   ├── assembler   负责将内部领域模型转化为可对外的DTO
│   └── cqe Command、Query和Event --  入参
│   └── dto Application层的所有接口返回值为DTO -- 出参
│   └── service 负责业务流程的编排，但本身不负责任何业务逻辑
├── domain
│   ├── aggregate 聚合
│   ├── entity 实体
│   ├── event 事件
│   │   ├── publish
│   │   └── subsctibe
│   ├── repo 接口
│   │   └── specification 统一封装查询
│   ├── service 领域服务
│   └── vo 值对象
└── infrastructure
│   ├── config 配置文件
│   ├── pkg 常用工具类封装（DB,log,tool等）
│   └── repository
│   ├── converter domain内对象转化 do {互转}
│   └── do 数据库映射对象
└── types 封装自定义的参数类型,例如 phone 自校验参数        
```

# 封装 zap 日志注入 trace 信息 Trace Id（内含 gin 例子）

### [hlog](https://github.com/hwholiday/learning_tools/tree/master/hlog) (源码地址)

- 实现自动切割文件 (基于 lumberjack 实现)
- 实现可传递 trace 信息 （基于 Context 实现）

```base
{"level":"info","ts":1639453661.4718382,"caller":"example/main.go:36","msg":"hconf example success"}
{"level":"info","ts":1639453664.7402327,"caller":"example/main.go:19","msg":"AddTraceId success","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
{"level":"info","ts":1639453664.7402515,"caller":"example/main.go:32","msg":"test","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
{"level":"debug","ts":1639453664.7402549,"caller":"example/main.go:33","msg":"test","traceId":"68867b89-c949-45a4-b325-86866c9f869a"}
```

### [hconfig 插拔式配置读取工具可动态加载](https://github.com/hwholiday/learning_tools/tree/master/hconfig)
- 支持 etcd
- 支持 kubernetes
- 支持 apollo
#### [使用文档](https://www.hwholiday.com/2022/hconfig/)
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


#### hconfig  使用

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

# go_push 一个实用的消息推送服务

### [go_push](https://github.com/hwholiday/learning_tools/tree/master/go_push) (推送服务)

    ```base
    ├── gateway // 长连接网关服务器
    │   ├── push_job.go    // 分发任务
    │   ├── room.go        // 房间，可作为某一类型的推送管理中心
    │   ├── room_manage.go // 房间管理
    │   ├── ws_conn.go     // 简单封装的websocket方法
    │   ├── ws_handle.go   // 处理websocket协议方法
    │   └── ws_server.go   // websocket服务
    ├── logic  //逻辑服务器    
    │   ├── http_handle.go // 推送，房间相关
    │   └── http_server.go // http服务
    └── main.go
    ```

### [HConf (基于etcd与viper的高可用配置中心)](https://github.com/hwholiday/learning_tools/tree/master/hconf)

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

### [HEvent (基于channel )](https://github.com/hwholiday/learning_tools/tree/master/hevent)
    1: 基于channel的简单事件订阅发布

### [设计模式](https://github.com/hwholiday/learning_tools/tree/master/design-patterns)
    1: 并发设计模式之Active Object
    2: Go 实现的责任链模式

### [micro_agent](https://github.com/hwholiday/learning_tools/tree/master/micro_agent) (micro微服务)

    1: base 基础方法
    2: conf 配置文件
    3：handler 对外处理方法
    4：model 数据格式
    5：proto protobuf 文件

### [all_packaged_library](https://github.com/hwholiday/learning_tools/tree/master/all_packaged_library) 里面封装了一些常用的库，有详细的介绍，持续更新

    1: base 里面封装mysql，redis，mgo，minio文件储存库S3协议，雪花算法，退出程序方法，redis全局锁，日志库等（插件形式可单独引用）
    2: logtool uber_zap日志库封装，可自动切分日志文件，压缩文件
    3: perf ppoof小插件
    4: push 集成苹果推送，google推送，华为推送
    5: quit 优雅的退出程序
    6: registrySelector 基于etcd实现的服务注册，发现，负载均衡

### [docker](https://github.com/hwholiday/learning_tools/tree/master/docker) (为你的服务插上docker_compose翅膀)

     1: docker 为你的服务插上docker_compose翅膀   

### [kafka](https://github.com/hwholiday/learning_tools/tree/master/kafka) (分布式消息发布订阅系统)

    1: main 消息队列

### [NATS_streaming](https://github.com/hwholiday/learning_tools/tree/master/NATS_streaming) (分布式消息发布订阅系统--是由NATS驱动的数据流系统)

    1: main 消息队列

### [nsq](https://github.com/hwholiday/learning_tools/tree/master/nsq) (分布式实时消息平台)

    1: main 消息队列

### [grpc](https://github.com/hwholiday/learning_tools/tree/master/grpc) (grpc学习)

    1: bidirectional_streaming_rpc 双向流grpc
    2: server_side_streaming_rpc 服务端流grpc，也可以写成客户端流grpc
    3: simple_rpc 简单grpc

### [rpc](https://github.com/hwholiday/learning_tools/tree/master/rpc) (rpc学习)

    1: main rpc学习

### [prometheus](https://github.com/hwholiday/learning_tools/tree/master/prometheus) (监控报警系统)

    1: server Prometheus监控报警系统

### [jaeger](https://github.com/hwholiday/learning_tools/tree/master/jaeger) (jaeger分布式链路追踪)

    1: main jaeger分布式链路追踪

### [service_load_balancing](https://github.com/hwholiday/learning_tools/tree/master/service_load_balancing) (负载均衡)

    1: fisher_yates_test  添加fisher-yates算法 负载均衡节点

### [hystrix](https://github.com/hwholiday/learning_tools/tree/master/hystrix) (熔断)

    1: hystrix 学习并使用熔断（Hystrix）

### [req_limit](https://github.com/hwholiday/learning_tools/tree/master/req_limit) (限流)

    1: main 使用带缓存的channel实现限流
    2: uber_ratelimit 使用uber_ratelimit实现限流       

### [ini](https://github.com/hwholiday/learning_tools/tree/master/ini) (配置文件库)

    1: main 配置文件ini的读取，以及自动匹配到结构体里面

### [minio](https://github.com/hwholiday/learning_tools/tree/master/minio) (对象存储服务)

    1: minio 对象存储服务使用

### [mysql](https://github.com/hwholiday/learning_tools/tree/master/mysql) (mysql服务器)

    1: main 简单的mysql使用

### [redis](https://github.com/hwholiday/learning_tools/tree/master/redis) (redis相关)

    1: bloom_filter redis 实现BloomFilter过滤器
    2: lock redis实现全局锁
    3: pipeline redis事务
    4: subscription redis发布订阅    

### [mongodb](https://github.com/hwholiday/learning_tools/tree/master/mongodb) (mongodb服务器)

    1: mgo.v2  mgo.v2库的基础使用学习
    2: mongo-go-driver  官方库的demo，以及事务提交（不能是单节点）

### [gin](https://github.com/hwholiday/learning_tools/tree/master/gin) (web框架gin学习)

    1: mvc 模式，swagger文档 可作为基础学习gin

### [jwt](https://github.com/hwholiday/learning_tools/tree/master/jwt) (JSON WEB TOKEN)

    1: jwt 学习使用   

### [snow_flake](https://github.com/hwholiday/learning_tools/tree/master/snow_flake) (雪花算法)

    1: main 雪花算法    

### [encryption_algorithm](https://github.com/hwholiday/learning_tools/tree/master/encryption_algorithm) (双棘轮算法, KDF链,迪菲-赫尔曼棘轮,x3dh)

    1: aes ase-(cfb,cbc,ecb)-(128,192,256)加解密方法
    2: curve25519 椭圆曲线算法
    3: 3curve25519 双棘轮算法,KDF链

### [LRU](https://github.com/hwholiday/learning_tools/tree/master/LRU) (缓存淘汰算法)

    1: list lru 缓存淘汰算法 

### [tcp](https://github.com/hwholiday/learning_tools/tree/master/tcp) (tcp协议实现)

    1: 实现网络库，封包，解包 len+tag+data 模式

### [websocket](https://github.com/hwholiday/learning_tools/tree/master/websocket) (websocket协议实现)

    1: 实现网络库

### [binary_conversion](https://github.com/hwholiday/learning_tools/tree/master/tool/binary_conversion) (进制转换)

    1: 10to36 10进制转36进制
    2: 10to62 10进制转62进制
    3: 10to76 10进制转76进制
    4: binary 用一个int64来描述开关（用户设置很多建议使用）       

### [job_worker_mode](https://github.com/hwholiday/learning_tools/tree/master/job_worker_mode) (job_worker模式)

    1: worker job_worker模式，可提高系统吞吐量

### [filewatch](https://github.com/hwholiday/learning_tools/tree/master/filewatch) (监控文件变化)

    1: main 监控文件变化 可实现自动构建

### [prometheus](https://github.com/hwholiday/learning_tools/tree/master/prometheus) (普罗米修斯)

    1: main 普罗米修斯    

### [goquery](https://github.com/hwholiday/learning_tools/tree/master/goquery) (网页解析工具)

    1: main 可以作为爬虫解析网页使用

### [active_object](https://github.com/hwholiday/learning_tools/tree/master/active_object) (并发设计模式)

    1: active_object Go并发设计模式之Active Object

### [heap](https://github.com/hwholiday/learning_tools/tree/master/container/heap) (优先级队列)

    1: heap 利用heap创建一个优先级队列

### [cli](https://github.com/hwholiday/learning_tools/tree/master/cli) (go命令行交互)

    1: main go命令行交互

### [context](https://github.com/hwholiday/learning_tools/tree/master/context) (context包学习)

    1: main context包学习 

### [err](https://github.com/hwholiday/learning_tools/tree/master/err) (error 相关)

    1: main golang 1.13 error 相关

### [interface](https://github.com/hwholiday/learning_tools/tree/master/interface) (interface包学习)

    1: main interface包学习
    2: middleware Golang 基于interface 实现中间件

### [syncPool](https://github.com/hwholiday/learning_tools/tree/master/syncPool) (syncPool包学习)

    1: main syncPool包学习

### [reflect](https://github.com/hwholiday/learning_tools/tree/master/reflect) (reflect包学习)

    1: main reflect包学习
