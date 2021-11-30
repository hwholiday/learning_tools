//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/adpter"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/aggregate"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/service"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/conf"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/mongo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/redis"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/repository"
)

//go:generate wire
var providerSet = wire.NewSet(
	conf.NewViper,
	conf.NewAppConfigCfg,
	conf.NewLoggerCfg,
	conf.NewRedisConfig,
	conf.NewMongoConfig,
	log.NewLogger,
	redis.NewRedis,
	mongo.NewMongo,
	repository.NewRepository,
	aggregate.NewFactory,
	service.NewService,
	adpter.NewSrv,
)

func NewApp() (*adpter.Server, error) {
	panic(wire.Build(providerSet))
}
