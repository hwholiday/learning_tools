package etcd

import (
	"github.com/hwholiday/learning_tools/hconfig/hconf"
	"go.etcd.io/etcd/api/v3/mvccpb"
)
import clientv3 "go.etcd.io/etcd/client/v3"

var _ hconf.DataWatcher = (*watcher)(nil)

type watcher struct {
	etcdConfig *etcdConfig
	ch         clientv3.WatchChan
	closeChan  chan struct{}
}

func newWatcher(s *etcdConfig) *watcher {
	w := &watcher{
		etcdConfig: s,
		ch:         nil,
		closeChan:  make(chan struct{}),
	}
	w.ch = s.client.Watch(s.options.ctx, s.options.root, clientv3.WithPrefix())
	return w
}

func (w *watcher) Change() ([]*hconf.Data, error) {
	select {
	case <-w.closeChan:
		return nil, nil
	case kv, ok := <-w.ch:
		if !ok {
			return nil, nil
		}
		var data []*mvccpb.KeyValue
		for _, v := range kv.Events {
			if v.Type == mvccpb.DELETE {
				continue
			}
			data = append(data, v.Kv)
		}
		return w.etcdConfig.kvsToData(data)
	}
}

func (w *watcher) Close() error {
	close(w.closeChan)
	return nil
}
