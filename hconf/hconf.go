package hconf

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type Register struct {
	etcdCli      *clientv3.Client
	opts         *Options
	mu           sync.Mutex
	confDef      map[string]*struct{}
	WatchConfKey []string `json:"watch_conf_key"`
}

func NewHConf(opt ...RegisterOptions) (*Register, error) {
	s := &Register{
		opts: newOptions(opt...),
	}
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	go s.watch()
	return s, nil
}

func (r *Register) PutConfKeyStruct(key string, val *struct{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	nodeKey := fmt.Sprintf("%s/%s", r.opts.RootName, key)
	r.WatchConfKey = append(r.WatchConfKey, nodeKey)
	r.confDef[nodeKey] = val
	return nil
}

func (r *Register) watch() {
	for {
		nodeKey := fmt.Sprintf("%s/%s", r.opts.RootName, "hw")
		time.Sleep(time.Second * 10)
		r.mu.Lock()
		val := r.confDef[nodeKey]
		r.mu.Unlock()

	}
}
