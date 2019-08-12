package config

type rdsConfig interface {
	GetIP() string
	GetPort() string
	GetPass() string
	GetMaxOpen() int
}

type defaultRedisConfig struct {
	IP      string `ini:"ip"`
	Port    string `ini:"port"`
	Pass    string `ini:"pass"`
	MaxOpen int    `ini:"max_open"`
}

func (m defaultRedisConfig) GetIP() string {
	return m.IP
}
func (m defaultRedisConfig) GetPort() string {
	return m.Port
}
func (m defaultRedisConfig) GetPass() string {
	return m.Pass
}
func (m defaultRedisConfig) GetMaxOpen() int {
	return m.MaxOpen
}
