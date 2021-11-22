package hconf

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
	"log"
	"sync"
	"time"
)

type Register interface {
	ConfByKey(key string, val interface{}) error
	Run() error
	Close() error
}

type register struct {
	etcdCli    *clientv3.Client
	opts       *Options
	viper      *viper.Viper
	mu         sync.Mutex
	confDef    map[string]interface{}
	needPutKey []string
}

func NewHConf(opt ...RegisterOptions) (Register, error) {
	s := &register{
		opts:    newOptions(opt...),
		confDef: make(map[string]interface{}),
		viper:   viper.New(),
	}
	s.viper.SetConfigFile(s.opts.LocalConfName)
	if !s.opts.UseLocalConf() {
		if len(s.opts.WatchRootName) == 0 {
			return nil, errors.New("watch_conf_key is empty")
		}
	}
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	if !s.checkEtcd() {
		log.Println("Link etcd failed to use local configuration")
		s.opts.UseLocal = true
	}
	return s, nil
}

func (r *register) Run() error {
	if r.opts.UseLocalConf() {
		return r.loadLocal()
	}
	if err := r.putNeedKv(); err != nil {
		return err
	}
	r.watchRookKey()
	return nil
}

func (r *register) putNeedKv() error {
	if len(r.needPutKey) == 0 {
		return nil
	}
	if err := r.viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range r.needPutKey {
		var mapData = make(map[string]interface{})
		if err := r.viper.UnmarshalKey(k, &mapData); err != nil {
			return err
		}
		if err := r.putEtcdConf(k, mapData); err != nil {
			return err
		}
		_ = r.viper.UnmarshalKey(k, r.confDef[k])
	}
	return nil
}

func (r *register) ConfByKey(key string, val interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.confDef[key] = val
	if r.opts.UseLocalConf() {
		return nil
	}
	res, err := r.readEtcdConf(key)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		r.needPutKey = append(r.needPutKey, key)
		return nil
	}
	return yaml.Unmarshal(res, val)
}

func (r *register) Close() error {
	if err := r.etcdCli.Close(); err != nil {
		return err
	}
	for k, _ := range r.confDef {
		r.viper.Set(k, r.confDef[k])
	}
	return r.viper.WriteConfig()
}

func (r *register) readEtcdConf(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opts.EtcdReadTimeOut)
	defer cancel()
	res, err := r.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	for _, v := range res.Kvs {
		if string(v.Key) == key {
			return v.Value, nil
		}
	}
	return nil, err
}

func (r *register) putEtcdConf(key string, val interface{}) error {
	if fmt.Sprint(val) == "map[]" {
		log.Println(fmt.Sprintf("No %s configuration found in etcd or local", key))
		return errors.New(fmt.Sprintf("No %s configuration found in etcd or local", key))
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.opts.EtcdReadTimeOut)
	defer cancel()
	data, err := yaml.Marshal(val)
	if err != nil {
		return err
	}
	_, err = r.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}
	log.Printf("put key %s ,data %s \n", key, string(data))
	return nil
}

func (r *register) loadLocalKey() error {
	if err := r.viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (r *register) loadLocal() error {
	if err := r.viper.ReadInConfig(); err != nil {
		return err
	}
	for k, _ := range r.confDef {
		if err := r.viper.UnmarshalKey(k, r.confDef[k]); err != nil {
			return err
		}
	}
	return nil
}

func (r *register) checkEtcd() bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.opts.EtcdReadTimeOut*time.Duration(len(r.opts.EtcdConf.Endpoints)))
	defer cancel()
	for _, v := range r.opts.EtcdConf.Endpoints {
		res, err := r.etcdCli.Status(ctx, v)
		if res == nil || err != nil {
			return false
		}
	}
	return true
}

func (r *register) watchRookKey() {
	if r.opts.UseLocalConf() {
		return
	}
	for _, v := range r.opts.WatchRootName {
		go func(reg *register, key string) {
			rch := r.etcdCli.Watch(context.Background(), key, clientv3.WithPrefix())
			for res := range rch {
				for _, ev := range res.Events {
					switch ev.Type {
					case mvccpb.PUT: //新增或修改
						k := string(ev.Kv.Key)
						if val, ok := reg.confDef[k]; ok {
							reg.mu.Lock()
							if err := yaml.Unmarshal(ev.Kv.Value, val); err != nil {
								log.Printf("change key %s ,data %s,err %s \n", k, string(ev.Kv.Value), err)
							}
							log.Printf("change key %s success data {%s}\n", k, string(ev.Kv.Value))
							reg.mu.Unlock()
						}
					}
				}
			}
		}(r, v)
	}

}
