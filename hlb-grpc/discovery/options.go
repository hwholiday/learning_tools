package discovery

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

var ErrLoadBalancingPolicy = errors.New("LoadBalancingPolicy is empty or not supported")
var ErrNoMatchFoundConn = errors.New("no match found conn")
var ErrNotGetRightConn = errors.New("did not get the right conn")

type Options struct {
	EtcdConf            clientv3.Config `json:"-"`
	SrvName             string
	LoadBalancingPolicy string
	Version             string
}

type ClientOptions func(*Options)

func SetName(name string) ClientOptions {
	return func(option *Options) {
		option.SrvName = fmt.Sprintf("/%s", strings.Replace(name, ".", "/", -1))
	}
}

func SetVersion(version string) ClientOptions {
	return func(option *Options) {
		option.Version = version
	}
}

func SetEtcdConf(conf clientv3.Config) ClientOptions {
	return func(options *Options) {
		options.EtcdConf = conf
	}
}

func SetLoadBalancingPolicy(name string) ClientOptions {
	return func(options *Options) {
		options.LoadBalancingPolicy = name
	}
}

func newOptions(opt ...ClientOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		},
		Version: "latest",
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func BuildCtxFilter(ctx context.Context, data map[string]string) context.Context {
	ctx = context.WithValue(ctx, customizeCtx, data)
	return ctx
}

func getCtxFilter(ctx context.Context) map[string]string {
	if ctx.Value(customizeCtx) == nil {
		return map[string]string{}
	}
	return ctx.Value(customizeCtx).(map[string]string)
}
