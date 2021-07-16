package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

type Register struct {
	etcdCli       *clientv3.Client
	leaseID       clientv3.LeaseID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	weight        string
}

func NewRegister(etcdCli *clientv3.Client, addr, weigit string, ttl int64) (*Register, error) {
	s := &Register{
		etcdCli: etcdCli,
		key:     "123123123123123123123",
		weight:  weigit,
	}
	resp, err := etcdCli.Grant(context.Background(), ttl)
	if err != nil {
		return s, err
	}
	_, err = etcdCli.Put(context.Background(), s.key, s.weight, clientv3.WithLease(resp.ID))
	if err != nil {
		return s, err
	}
	leaseRespChan, err := etcdCli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return s, err
	}
	s.leaseID = resp.ID
	s.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s  weight:%s  success!", s.key, s.weight)
	return s, nil
}

func (s *Register) ListenLeaseRespChan() (isClose bool) {
	for range s.keepAliveChan {
	}
	return true
}

// Close 注销服务
func (s *Register) Close() error {
	if _, err := s.etcdCli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Println("撤销租约")
	return s.etcdCli.Close()
}
