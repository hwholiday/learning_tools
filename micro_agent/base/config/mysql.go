package config

type sqlConfig interface {
	GetIp() string
	GetPort() string
	GetName() string
	GetPass() string
	GetDb() string
	GetMaxIdle() int
	GetMaxOpen() int
}

type defaultMysqlConfig struct {
	Ip      string `ini:"ip"`
	Port    string `ini:"port"`
	Name    string `ini:"name"`
	Pass    string `ini:"pass"`
	Db      string `ini:"db"`
	MaxIdle int    `ini:"max_idle"`
	MaxOpen int    `ini:"max_open"`
}

func (m defaultMysqlConfig) GetIp() string {
	return m.Ip
}
func (m defaultMysqlConfig) GetPort() string {
	return m.Port
}
func (m defaultMysqlConfig) GetName() string {
	return m.Name
}
func (m defaultMysqlConfig) GetPass() string {
	return m.Pass
}
func (m defaultMysqlConfig) GetDb() string {
	return m.Db
}
func (m defaultMysqlConfig) GetMaxIdle() int {
	return m.MaxIdle
}
func (m defaultMysqlConfig) GetMaxOpen() int {
	return m.MaxOpen
}
