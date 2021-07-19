package discovery

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Options struct {
	EtcdConf clientv3.Config `json:"-"`
	SrvName  string
}

type ClientOptions func(*Options)

func SetName(name string) ClientOptions {
	return func(option *Options) {
		option.SrvName = fmt.Sprintf("/%s", strings.Replace(name, ".", "/", -1))
	}
}

func SetEtcdConf(conf clientv3.Config) ClientOptions {
	return func(options *Options) {
		options.EtcdConf = conf
	}
}

func newOptions(opt ...ClientOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		},
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}
