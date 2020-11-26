```base
├── api
├── cmd		项目启动文件
│   └── user		项目名称
│       └── main.go
├── conf		配置文件   
│   └── user.yaml
├── internal
│   └── user		项目名称
│       ├── adapter		适配器,对外提供不同的协议适配
│       │   ├── grpc
│       │   └── http
│       ├── domain
│       │   ├── aggregate		聚合根
│       │   ├── dto						定义入参,出参
│       │   ├── entity				实体
│       │   ├── event					事件
│       │   ├── interface		抽象接口由外部 repository 实现
│       │   ├── service				领域服务
│       │   └── valobj				值对象
│       ├── facade				接口防腐层 调用其他服务接口
│       │
│       └── infrastructure		适配器，是对技术实现的适配
│           └── repository					依赖倒置实现储存服务
└── pkg   三方库工具类


```

#### 提交空目录
``` based
 //项目的根目录下，执行下面语句：
 find . -name .git -prune -o -type d -empty -exec touch {}/.gitignore \; 
```