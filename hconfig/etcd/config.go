package etcd

import (
	"errors"
	"fmt"
	"github.com/hwholiday/learning_tools/hconfig/hconf"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
)

var _ hconf.DataSource = (*etcdConfig)(nil)

type etcdConfig struct {
	client  *clientv3.Client
	options *options
}

func NewEtcdConfig(cli *clientv3.Client, opts ...Option) (hconf.DataSource, error) {
	if cli == nil {
		return nil, errors.New("etcd client is nil")
	}
	conf := &etcdConfig{
		client:  cli,
		options: NewOptions(opts...),
	}
	return conf, nil
}

func (c *etcdConfig) Load() ([]*hconf.Data, error) {
	data := make([]*hconf.Data, 0)
	for _, v := range c.options.paths {
		loadData, err := c.loadPath(v)
		if err != nil {
			return nil, err
		}
		data = append(data, loadData)
	}
	return data, nil
}

func (c *etcdConfig) loadPath(path string) (*hconf.Data, error) {
	rsp, err := c.client.Get(c.options.ctx, fmt.Sprintf("%s/%s", c.options.root, path))
	if err != nil {
		return nil, err
	}
	data := new(hconf.Data)
	for _, item := range rsp.Kvs {
		k := string(item.Key)
		k = strings.ReplaceAll(strings.ReplaceAll(k, c.options.root, ""), "/", "")
		if k == path {
			data.Key = k
			data.Val = item.Value
			break
		}
	}
	return data, nil
}

func (c *etcdConfig) kvsToData(kvs []*mvccpb.KeyValue) ([]*hconf.Data, error) {
	data := make([]*hconf.Data, 0)
	for _, item := range kvs {
		k := string(item.Key)
		k = strings.ReplaceAll(strings.ReplaceAll(k, c.options.root, ""), "/", "")
		for _, v := range c.options.paths {
			if k == v {
				data = append(data, &hconf.Data{
					Key: k,
					Val: item.Value,
				})
			}
		}
	}
	return data, nil
}

func (c *etcdConfig) Watch() (hconf.DataWatcher, error) {
	return newWatcher(c), nil
}
