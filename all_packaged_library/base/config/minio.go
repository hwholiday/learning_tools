package config

type minConfig interface {
	GetPath() string
	GetAccessKeyId() string
	GetSecretAccessKey() string
}

type defaultMinioConfig struct {
	Path            string `ini:"path"`
	AccessKeyId     string `ini:"access_key_id"`
	SecretAccessKey string `ini:"secret_access_key"`
}

func (m defaultMinioConfig) GetPath() string {
	return m.Path
}

func (m defaultMinioConfig) GetAccessKeyId() string {
	return m.AccessKeyId
}

func (m defaultMinioConfig) GetSecretAccessKey() string {
	return m.SecretAccessKey
}
