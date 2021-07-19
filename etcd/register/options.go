package register

import (
	"fmt"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
	"reflect"
	"strings"
	"time"
)

type Node struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Address string `json:"address"`
}

type Options struct {
	EtcdConf    clientv3.Config      `json:"-"`
	RegisterTtl int64                `json:"-"`
	Node        *Node                `json:"node"`
	Metadata    map[string]string    `json:"metadata"`
	Endpoints   map[string]Endpoints `json:"endpoints"`
}

type Endpoints struct {
	Online         bool   `json:"online"`
	LimitingSwitch bool   `json:"limiting_switch"` //每秒多少次
	Limiting       int64  `json:"limiting"`        //每秒多少次
	Fuse           bool   `json:"fuse"`            //熔断
	Defaults       []byte `json:"defaults"`        //熔断默认值
}

type RegisterOptions func(*Options)

func newOptions(opt ...RegisterOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		},
		Node:        &Node{Version: "latest"},
		RegisterTtl: 10,
		Endpoints:   make(map[string]Endpoints),
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func SetName(name string) RegisterOptions {
	return func(options *Options) {
		path := strings.Split(name, ".")
		options.Node.Name = path[len(path)-1]
		options.Node.Id = fmt.Sprintf("%s-%s", options.Node.Name, uuid.Must(uuid.NewUUID()).String())
		options.Node.Path = fmt.Sprintf("/%s", strings.Join(path, "/"))
	}
}

func SetSrv(srv interface{}) RegisterOptions {
	return func(options *Options) {
		typ := reflect.TypeOf(srv)
		for m := 0; m < reflect.TypeOf(srv).NumMethod(); m++ {
			options.Endpoints[typ.Method(m).Name] = Endpoints{
				Online:         true,
				LimitingSwitch: false,
				Limiting:       10,
				Fuse:           false,
				Defaults:       nil,
			}
		}
	}
}

func SetRegisterTtl(registerTtl int64) RegisterOptions {
	return func(options *Options) {
		options.RegisterTtl = registerTtl
	}
}

func SetVersion(version string) RegisterOptions {
	return func(options *Options) {
		options.Node.Version = version
	}
}
func SetEtcdConf(conf clientv3.Config) RegisterOptions {
	return func(options *Options) {
		options.EtcdConf = conf
	}
}

func SetAddress(address string) RegisterOptions {
	return func(options *Options) {
		options.Node.Address = address
	}
}

func SetMetadata(metadata map[string]string) RegisterOptions {
	return func(options *Options) {
		options.Metadata = metadata
	}
}
