package config

import (
	"path/filepath"
	"sync"
)
import "github.com/go-ini/ini"

var (
	mysqlConfig defaultMysqlConfig
	serConfig   defaultServerConfig
	utilsConfig defaultLogToolConfig
	m           sync.Mutex
)

func Init(path string) {
	var (
		err error
		cfg *ini.File
	)

	m.Lock()
	defer m.Unlock()
	if err = ini.MapTo(&mysqlConfig, filepath.Join(path, "mysql.ini")); err != nil {
		panic(err)
	}
	if err = ini.MapTo(&serConfig, filepath.Join(path, "user_agent.ini")); err != nil {
		panic(err)
	}
	if cfg, err = ini.Load(filepath.Join(path, "tool.ini")); err != nil {

	}
	if err = cfg.Section("zap").MapTo(&utilsConfig); err != nil {
		panic(err)
	}
}

func GetMysqlConfig() (fig sqlConfig) {
	return mysqlConfig
}
func GetServerConfig() (fig serverConfig) {
	return serConfig
}
func GetToolLogConfig() (fig toolLogConfig) {
	return utilsConfig
}
