package config

type serversConfig interface {
	GetEtcd() string
	GetName() string
	GetAddr() string
}

type defaultServerConfig struct {
	Etcd string `ini:"etcd"`
	Name string `ini:"name"`
	Addr string `ini:"addr"`
}

func (m defaultServerConfig) GetEtcd() string {
	return m.Etcd
}

func (m defaultServerConfig) GetName() string {
	return m.Name
}
func (m defaultServerConfig) GetAddr() string {
	return m.Addr
}
