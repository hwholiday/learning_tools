package hconfig

import (
	"errors"
	"github.com/hwholiday/learning_tools/hconfig/hconf"
)

type Option func(opt *options)

type options struct {
	dataSource hconf.DataSource
	useLocal   bool
}

func WithDataSource(d hconf.DataSource) Option {
	return func(opt *options) {
		opt.dataSource = d
	}
}

func UseLocal() Option {
	return func(opt *options) {
		opt.useLocal = true
	}
}

func newOptions(opts ...Option) (*options, error) {
	o := &options{
		dataSource: nil,
		useLocal:   false,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.dataSource == nil {
		return nil, errors.New("dataSource is nil")
	}
	return o, nil
}
