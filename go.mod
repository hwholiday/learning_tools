module github.com/hwholiday/learning_tools

go 1.16

replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.27.0

require (
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/Shopify/sarama v1.30.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.1
	github.com/garyburd/redigo v1.6.3
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.64.0
	github.com/go-kit/kit v0.12.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/memberlist v0.3.0
	github.com/hpcloud/tail v1.0.0
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/stan.go v0.10.2
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nsqio/go-nsq v1.1.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.11.0
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.3.1
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/viper v1.9.0
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.4
	github.com/tsuna/gohbase v0.0.0-20211118233222-1c6789fac7d4
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/urfave/cli v1.22.5
	github.com/urfave/cli/v2 v2.3.0
	go.etcd.io/etcd v3.3.27+incompatible
	go.etcd.io/etcd/api/v3 v3.5.1
	go.etcd.io/etcd/client/v3 v3.5.1
	go.mongodb.org/mongo-driver v1.7.4
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/nats-io/nats-streaming-server v0.23.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
)
