package hconf

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Options struct {
	EtcdConf      clientv3.Config `json:"-"`
	RegisterTtl   time.Duration   `json:"register_ttl"`
	LocalConfName string          `json:"local_conf_name"`
	UseLocal      bool            `json:"use_local"`
	WatchRoot     []string        `json:"watch_root"`
}

type RegisterOptions func(*Options)

func newOptions(opt ...RegisterOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 3 * time.Second,
		},
		RegisterTtl:   2 * time.Second,
		LocalConfName: "./hconf.yaml",
		UseLocal:      false,
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

func SetUseLocal() RegisterOptions {
	return func(options *Options) {
		options.UseLocal = true
	}
}

func SetLocalConfName(name string) RegisterOptions {
	return func(options *Options) {
		options.LocalConfName = name
	}
}

func SetWatchRoot(name []string) RegisterOptions {
	return func(options *Options) {
		options.WatchRoot = name
	}
}
