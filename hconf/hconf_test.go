package hconf

import (
	"testing"
)

type Conf struct {
	Net  Net
	Net2 Net
	Net3 Net
}

type Net struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func TestHConf(t *testing.T) {
	var conf = Conf{}
	r, err := NewHConf(
		SetWatchRootName([]string{"/gs/conf"}),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r.ConfByKey("/gs/conf/net", &conf.Net))
	t.Log(r.ConfByKey("/gs/conf/net2222", &conf.Net2))
	t.Log(r.ConfByKey("/gs/conf/net3333", &conf.Net3))
	if err := r.Run(); err != nil {
		t.Error(err)
		return
	}
	t.Log(conf)
	t.Log(r.Close())
}
