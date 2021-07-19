package register

import (
	"fmt"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	EtcdConf    clientv3.Config   `json:"-"`
	RegisterTtl int64             `json:"-"`
	Node        *Node             `json:"node"`
	Metadata    map[string]string `json:"metadata"`
	Endpoints   map[string]string `json:"endpoints"`
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
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func SetName(name string) RegisterOptions {
	return func(options *Options) {
		path := strings.Split(name, ".")
		if options.Node == nil {
			options.Node = &Node{}
		}
		options.Node.Name = path[len(path)-1]
		options.Node.Id = fmt.Sprintf("%s-%s", options.Node.Name, uuid.Must(uuid.NewUUID()).String())
		options.Node.Path = fmt.Sprintf("/%s", strings.Join(path, "/"))
	}
}

func SetRegisterTtl(registerTtl int64) RegisterOptions {
	return func(options *Options) {
		if options.Node == nil {
			options.Node = &Node{}
		}
		options.RegisterTtl = registerTtl
	}
}

func SetVersion(version string) RegisterOptions {
	return func(options *Options) {
		if options.Node == nil {
			options.Node = &Node{}
		}
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
		if options.Node == nil {
			options.Node = &Node{}
		}
		options.Node.Address = address
	}
}

func SetMetadata(metadata map[string]string) RegisterOptions {
	return func(options *Options) {
		options.Metadata = metadata
	}
}
