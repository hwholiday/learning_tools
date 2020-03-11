## learning_tools [源码地址](https://github.com/hwholiday/learning_tools)

# go-kit 微服务实践，从入门到精通系列    
### [go-kit](https://github.com/hwholiday/learning_tools/tree/master/go-kit) (go-kit微服务)
    1: v1 go-kit 微服务 基础使用 （HTTP）
    2: v2 go-kit 微服务 添加日志（user/zap ,并为每个请求添加UUID） 
    3: v3 go-kit 微服务 身份认证 （JWT）
    4: v4 go-kit 微服务 限流 （uber/ratelimit 和 golang/rate 实现）
    5: v5 go-kit 微服务 使用GRPC（并为每个请求添加UUID） 
    6: v6 go-kit 微服务 服务注册与发现（etcd实现）
    7: v7 go-kit 微服务 服务监控（prometheus 实现）
    8: v8 go-kit 微服务 服务熔断（hystrix-go 实现）
    9: v9 go-kit 微服务 服务链路追踪（jaeger 实现）(1)
    10: v9 go-kit 微服务 服务链路追踪（jaeger 实现）(2)
    11: v11 go-kit 微服务 添加一个简单网关

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

### [micro_agent](https://github.com/hwholiday/learning_tools/tree/master/micro_agent) (micro微服务)
    1: base 基础方法
    2: conf 配置文件
    3：handler 对外处理方法
    4：model 数据格式
    5：proto protobuf 文件

###  [all_packaged_library](https://github.com/hwholiday/learning_tools/tree/master/all_packaged_library) 里面封装了一些常用的库，有详细的介绍，持续更新
    1: base 里面封装mysql，redis，mgo，minio文件储存库S3协议，雪花算法，退出程序方法，redis全局锁，日志库等（插件形式可单独引用）
    2: logtool uber_zap日志库封装，可自动切分日志文件，压缩文件
    3: perf ppoof小插件
    4: push 集成苹果推送，google推送，华为推送
    5: quit 优雅的退出程序
    6: registrySelector 基于etcd实现的服务注册，发现，负载均衡

### [nsq](https://github.com/hwholiday/learning_tools/tree/master/docker) (为你的服务插上docker_compose翅膀)
     1: docker 为你的服务插上docker_compose翅膀   
   
### [kafka](https://github.com/hwholiday/learning_tools/tree/master/kafka) (分布式消息发布订阅系统)
    1: main 消息队列
    
### [NATS_streaming](https://github.com/hwholiday/learning_tools/tree/master/NATS_streaming)(分布式消息发布订阅系统--是由NATS驱动的数据流系统)
    1: main 消息队列
    
### [nsq](https://github.com/hwholiday/learning_tools/tree/master/kafka) (分布式实时消息平台)
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

### [service_registration_discovery](https://github.com/hwholiday/learning_tools/tree/master/service_registration_discovery) (服务注册与发现)
    1: etcdv3  通过etcd实现服务注册与发现

### [service_req_hystrix](https://github.com/hwholiday/learning_tools/tree/master/service_req_hystrix) (熔断)
    1: hystrix 学习并使用熔断（Hystrix）

### [service_req_limit](https://github.com/hwholiday/learning_tools/tree/master/service_req_limit) (限流)
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
    
### [zaplog](https://github.com/hwholiday/learning_tools/tree/master/kafka) (uber zap日志封装)
    1: zap 日志封装    

### [snow_flake](https://github.com/hwholiday/learning_tools/tree/master/snow_flake) (雪花算法)
    1: main 雪花算法    

### [encryption_algorithm](https://github.com/hwholiday/learning_tools/tree/master/encryption_algorithm) (双棘轮算法, KDF链,迪菲-赫尔曼棘轮,x3dh)
    1: aes ase-(cfb,cbc,ecb)-(128,192,256)加解密方法
    2: curve25519 椭圆曲线算法
    3: 3curve25519 双棘轮算法,KDF链

### [LRU](https://github.com/hwholiday/learning_tools/tree/master/LRU)(缓存淘汰算法)
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

### [push](https://github.com/hwholiday/learning_tools/tree/master/prometheus) (一个简单的推送服务)
    1: main 推送服务    
    
### [goquery](https://github.com/hwholiday/learning_tools/tree/master/goquery) (网页解析工具)
    1: main 可以作为爬虫解析网页使用
 
### [active_object](https://github.com/hwholiday/learning_tools/tree/master/active_object)(并发设计模式)
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