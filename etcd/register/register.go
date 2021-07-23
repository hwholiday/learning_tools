package register

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Register struct {
	etcdCli       *clientv3.Client
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	opts          *Options
	name          string
}

func NewRegister(opt ...RegisterOptions) (*Register, error) {
	s := &Register{
		opts: newOptions(opt...),
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Duration(s.opts.RegisterTtl)*time.Second)
	defer cancel()
	data, err := json.Marshal(s.opts)
	if err != nil {
		return nil, err
	}
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	//申请租约
	resp, err := etcdCli.Grant(ctx, s.opts.RegisterTtl)
	if err != nil {
		return s, err
	}
	s.name = fmt.Sprintf("%s/%s", s.opts.Node.Path, s.opts.Node.Id)
	//注册节点
	_, err = etcdCli.Put(ctx, s.name, string(data), clientv3.WithLease(resp.ID))
	if err != nil {
		return s, err
	}
	//续约租约
	s.keepAliveChan, err = etcdCli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s *Register) ListenKeepAliveChan() (isClose bool) {
	for range s.keepAliveChan {
	}
	return true
}

// Close 注销服务
func (s *Register) Close() error {
	if _, err := s.etcdCli.Delete(context.Background(), s.name); err != nil {
		return err
	}
	return s.etcdCli.Close()
}
