package config

import (
	"path/filepath"
	"sync"
)
import "github.com/go-ini/ini"

var (
	utilsConfig  defaultLogToolConfig
	redisConfig  defaultRedisConfig
	minioConfig  defaultMinioConfig
	mgoConfig    defaultMgoConfig
	serverConfig defaultServerConfig
	mysqlConfig  defaultMysqlConfig

	m            sync.Mutex
)

func Init(path string) {
	var (
		err error
		cfg *ini.File
	)

	m.Lock()
	defer m.Unlock()
	if cfg, err = ini.Load(filepath.Join(path, "db.ini")); err != nil {
		panic(err)
	}
	if err = cfg.Section("redis").MapTo(&redisConfig); err != nil {
		panic(err)
	}
	if err = cfg.Section("mysql").MapTo(&mysqlConfig); err != nil {
		panic(err)
	}
	if err = cfg.Section("mongodb").MapTo(&mgoConfig); err != nil {
		panic(err)
	}
	if err = cfg.Section("minio").MapTo(&minioConfig); err != nil {
		panic(err)
	}
	if cfg, err = ini.Load(filepath.Join(path, "tool.ini")); err != nil {
		panic(err)
	}
	if err = cfg.Section("zap").MapTo(&utilsConfig); err != nil {
		panic(err)
	}
	if err = cfg.Section("server").MapTo(&serverConfig); err != nil {
		panic(err)
	}
}

func GetToolLogConfig() (fig toolLogConfig) {
	return utilsConfig
}

func GetRedisConfig() (fig rdsConfig) {
	return redisConfig
}

func GetMgoConfig() (fig mgConfig) {
	return mgoConfig
}
func GetServerConfig() (fig serversConfig) {
	return serverConfig
}

func GetMinioConfig() (fig minConfig) {
	return minioConfig
}

func GetMysqlConfig() msqlConfig {
	return mysqlConfig
}

