### Golang DDD 的项目分层结构（六边形架构）

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