package config

type mgConfig interface {
	GetUrl() string
	GetName() string
	GetPass() string
}

type defaultMgoConfig struct {
	Url  string `ini:"url"`
	Name string `ini:"name"`
	Pass string `ini:"pass"`
}

func (m defaultMgoConfig) GetUrl() string {
	return m.Url
}
func (m defaultMgoConfig) GetName() string {
	return m.Name
}
func (m defaultMgoConfig) GetPass() string {
	return m.Pass
}
