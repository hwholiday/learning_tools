package config

type serverConfig interface {
	GetAppMode()string
	AppIsDebug()bool
	GetServerName()string
	GetEtcdAddr()string
	GetBlack()[]string
}

type defaultServerConfig struct {
	 AppMode string `ini:"app_mode"`
	 ServerName string `ini:"server_name"`
	 EtcdAddr string `ini:"etcd_addr"`
	 List  list `ini:"list"`
}
type list struct {
	Black []string `ini:"black"`
}

func (s defaultServerConfig)GetAppMode()string {
	return s.AppMode
}
func (s defaultServerConfig)AppIsDebug()bool {
	if s.AppMode=="debug"{
		return true
	}
	return false
}
func (s defaultServerConfig)GetServerName()string {
	return s.ServerName
}
func (s defaultServerConfig)GetEtcdAddr()string {
	return s.EtcdAddr
}
func (s defaultServerConfig)GetBlack()[]string {
	return s.List.Black
}