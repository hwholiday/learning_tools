package hconf

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Options struct {
	EtcdConf    clientv3.Config `json:"-"`
	RegisterTtl int64           `json:"-"`
	RootName    string          `json:"root_Name"`
	UseLocal    bool            `json:"use_local"`
}

type RegisterOptions func(*Options)

func newOptions(opt ...RegisterOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		},
		UseLocal: false,
		RootName: "hconf",
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func SetEtcdConf(conf clientv3.Config) RegisterOptions {
	return func(options *Options) {
		options.EtcdConf = conf
	}
}

func SetName(name string) RegisterOptions {
	return func(options *Options) {
		options.RootName = name
	}
}

func UseLocal() RegisterOptions {
	return func(options *Options) {
		options.UseLocal = true
	}
}
