package kubernetes

import (
	"golang.org/x/net/context"
)

type Option func(*options)

// /honf/test 对应一个configMap
// /honf/test2 对应一个configMap
// /honf/test3 对应一个configMap
// /honf/test3/conf/2 对应一个configMap
type options struct {
	ctx            context.Context
	namespace      string
	paths          []string
	kubeConfigPath string
	masterUrl      string
}

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithNamespace(namespace string) Option {
	return func(o *options) {
		o.namespace = namespace
	}
}

func WithPaths(path ...string) Option {
	return func(o *options) {
		o.paths = path
	}
}

func KubeConfigPath(kubeConfigPath string) Option {
	return func(o *options) {
		o.kubeConfigPath = kubeConfigPath
	}
}

func KubeMasterUrl(masterUrl string) Option {
	return func(o *options) {
		o.masterUrl = masterUrl
	}
}

func NewOptions(opts ...Option) *options {
	options := &options{
		ctx:       context.Background(),
		namespace: GetNamespace(),
		paths:     []string{"test"},
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
