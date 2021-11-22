package hconf

import (
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type Register struct {
	etcdCli      *clientv3.Client
	opts         *Options
	mu           sync.Mutex
	confDef      map[string]interface{}
	WatchConfKey []string `json:"watch_conf_key"`
}

func NewHConf(opt ...RegisterOptions) (*Register, error) {
	s := &Register{
		opts:    newOptions(opt...),
		confDef: make(map[string]interface{}),
	}
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	return s, nil
}

func (r *Register) Run() {
	r.loadLocal()
}

func (r *Register) loadLocal() {
	v := viper.New()
	v.SetConfigFile(r.opts.LocalConfName)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("1111", err)
		return
	}
	for k, _ := range r.confDef {
		if err := v.UnmarshalKey(k, r.confDef[k]); err != nil {
			fmt.Println("22222", k, err)
			return
		}
	}
}

func (r *Register) GetConfByKey(key string, val interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.confDef[key] = val
	if r.opts.UseLocal {
		return
	}
}
