### 仿微信 auth2 授权登陆
##### DDD（领域设计驱动）+六边形架构

```base
.
├── adpter
│   ├── adpter.go
│   └── http
│       ├── auth_handles
│       │   ├── auth_code_handles.go
│       │   ├── auth_token_handles.go
│       │   └── handers.go
│       ├── http.go
│       └── routers
│           ├── middleware.go
│           └── routers.go
├── cmd
│   ├── app.yaml
│   ├── cmd
│   ├── main.go
│   ├── wire_gen.go
│   └── wire.go
├── domain
│   ├── aggregate
│   │   ├── auth_code.go
│   │   ├── auth_factory.go
│   │   ├── auth_token.go
│   │   ├── auth_token_produce.go
│   │   └── factory.go
│   ├── dto
│   │   ├── auth_code.go
│   │   ├── auth_token.go
│   │   └── user.go
│   ├── entity
│   │   └── merchant.go
│   ├── obj
│   │   ├── auth_code.go
│   │   └── auth_token.go
│   ├── repo
│   │   ├── auth_code.go
│   │   ├── auth_token.go
│   │   ├── merchant.go
│   │   └── specification
│   │       ├── auth_code_by_code.go
│   │       ├── auth_token_by_code.go
│   │       └── merchant_by_appid.go
│   └── service
│       ├── auth_code.go
│       ├── auth_token.go
│       ├── merchant.go
│       └── service.go
├── infrastructure
│   ├── conf
│   │   ├── auth_consts.go
│   │   └── conf.go
│   ├── pkg
│   │   ├── database
│   │   │   ├── mongo
│   │   │   │   └── mgo.go
│   │   │   └── redis
│   │   │       ├── lock.go
│   │   │       ├── redis.go
│   │   ├── hcode
│   │   │   ├── base.go
│   │   │   └── code.go
│   │   ├── log
│   │   │   └── zap.go
│   │   └── tool
│   │       ├── aes.go
│   │       ├── aes_test.go
│   │       ├── jwt.go
│   │       └── jwt_test.go
│   └── repository
│       ├── atuh_code.go
│       ├── atuh_token.go
│       ├── merchant.go
│       └── repository.go
└── README.md

```


#### APPID 必须是10位 因为使用了 ase126 来作为 OpenId的生成

##### key=APPID(10)+SALT(6) 16 = 126