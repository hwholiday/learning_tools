package hconf

import "testing"

type Conf struct {
	Net Net
}

type Net struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func TestHConf(t *testing.T) {
	var conf = Conf{}
	r, err := NewHConf()
	if err != nil {
		t.Error(err)
		return
	}
	r.GetConfByKey("/gs/conf/net", &conf.Net)
	r.Run()
	t.Log(conf)
}
