package hconf

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Options struct {
	EtcdConf        clientv3.Config
	EtcdReadTimeOut time.Duration
	LocalConfName   string
	UseLocal        bool
	WatchRootName   []string
}

type RegisterOptions func(*Options)

func newOptions(opt ...RegisterOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 3 * time.Second,
		},
		EtcdReadTimeOut: 2 * time.Second,
		LocalConfName:   "./hconf.yaml",
		UseLocal:        false,
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func (p *Options) UseLocalConf() bool {
	return p.UseLocal
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
func SetEtcdReadTimeOut(t time.Duration) RegisterOptions {
	return func(options *Options) {
		options.EtcdReadTimeOut = t
	}
}

func SetLocalConfName(name string) RegisterOptions {
	return func(options *Options) {
		options.LocalConfName = name
	}
}

func SetWatchRootName(name []string) RegisterOptions {
	return func(options *Options) {
		options.WatchRootName = name
	}
}
