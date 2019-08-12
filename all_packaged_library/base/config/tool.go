package config

type toolLogConfig interface {
	GetDevelopment() bool
	GetLogFileDir() string
	GetAppName() string
	GetMaxSize() int
	GetMaxBackups() int
	GetMaxAge() int
	GetDcId() int64
}

type defaultLogToolConfig struct {
	Development string `ini:"development"`
	LogFileDir  string `ini:"logFileDir"`
	AppName     string `ini:"appName"`
	MaxSize     int    `ini:"maxSize"`
	MaxBackups  int    `ini:"maxBackups"`
	MaxAge      int    `ini:"maxAge"`
	Dc          int64  `ini:"dc"`
}

func (m defaultLogToolConfig) GetDevelopment() bool {
	if m.Development == "true" {
		return true
	} else {
		return false
	}
}
func (m defaultLogToolConfig) GetLogFileDir() string {
	return m.LogFileDir
}
func (m defaultLogToolConfig) GetAppName() string {
	return m.AppName
}
func (m defaultLogToolConfig) GetMaxSize() int {
	return m.MaxSize
}
func (m defaultLogToolConfig) GetMaxBackups() int {
	return m.MaxBackups
}
func (m defaultLogToolConfig) GetMaxAge() int {
	return m.MaxAge
}
func (m defaultLogToolConfig) GetDcId() int64 {
	return m.Dc
}
