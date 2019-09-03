package config

type msqlConfig interface {
	GetIp() string
	GetPort() int
	GetUser() string
	GetPwd() string
	GetDbName() string
	GetPoolSize() int
}

type defaultMysqlConfig struct {
	Ip       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Pwd      string `ini:"pwd"`
	DbName   string `ini:"dbName"`
	PoolSize int    `ini:"poolSize"`
}

func (m defaultMysqlConfig) GetIp() string {
	return m.Ip
}

func (m defaultMysqlConfig) GetPort() int {
	return m.Port
}

func (m defaultMysqlConfig) GetUser() string {
	return m.User
}

func (m defaultMysqlConfig) GetPwd() string {
	return m.Pwd
}
func (m defaultMysqlConfig) GetDbName() string {
	return m.DbName
}
func (m defaultMysqlConfig) GetPoolSize() int {
	return m.PoolSize
}
