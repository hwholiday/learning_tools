package conf

import (
	"flag"
	"fmt"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/mongo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/redis"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	NetConf *NetConf `yaml:"netConf"`
}

type NetConf struct {
	Name         string `yaml:"name"`         // 服务名称
	ServerAddr   string `yaml:"serverAddr"`   // 服务地址
	ReadTimeout  int    `yaml:"readTimeout"`  // 单位s
	WriteTimeout int    `yaml:"writeTimeout"` // 单位s
}

var conf = flag.String("conf", "./app.yaml", "conf")

func NewViper() (*viper.Viper, error) {
	flag.Parse()
	v := viper.New()
	v.SetConfigFile(*conf)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%w unmarshal ReadInConfig error", err)
	}
	return v, nil
}

func NewAppConfigCfg(v *viper.Viper) (conf *AppConfig, err error) {
	conf = new(AppConfig)
	if err = v.UnmarshalKey("netConf", &conf.NetConf); err != nil {
		err = fmt.Errorf("%w unmarshal log error", err)
	}
	return
}

func NewLoggerCfg(v *viper.Viper) (conf *log.Options, err error) {
	conf = new(log.Options)
	if err = v.UnmarshalKey("log", conf); err != nil {
		err = fmt.Errorf("%w unmarshal log error", err)
	}
	return
}

func NewMongoConfig(v *viper.Viper) (conf *mongo.Config, err error) {
	conf = new(mongo.Config)
	if err = v.UnmarshalKey("mongo", &conf); err != nil {
		err = fmt.Errorf("%w unmarshal NewMongoConfig error", err)
	}
	return
}

func NewRedisConfig(v *viper.Viper) (conf *redis.Config, err error) {
	conf = new(redis.Config)
	if err = v.UnmarshalKey("redis", &conf); err != nil {
		err = fmt.Errorf("%w unmarshal NewRedisConfig error", err)
	}
	return
}
