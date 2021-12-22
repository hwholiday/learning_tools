package kubernetes

import (
	"github.com/hwholiday/learning_tools/hconfig/hconf"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

var _ hconf.DataWatcher = (*watcher)(nil)

type watcher struct {
	kubernetesConfig *kubernetesConfig
	watch            watch.Interface
	closeChan        chan struct{}
}

func newWatcher(s *kubernetesConfig) (*watcher, error) {
	w := &watcher{
		kubernetesConfig: s,
		closeChan:        make(chan struct{}),
	}
	watch, err := s.client.CoreV1().ConfigMaps(s.options.namespace).Watch(s.options.ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	w.watch = watch
	return w, nil
}

func (w *watcher) Change() ([]*hconf.Data, error) {
	select {
	case <-w.closeChan:
		return nil, nil
	case kv, ok := <-w.watch.ResultChan():
		if !ok {
			return nil, nil
		}
		cm, ok := kv.Object.(*v1.ConfigMap)
		if !ok {
			return nil, nil
		}
		if kv.Type == "DELETED" {
			return nil, nil
		}
		return w.kubernetesConfig.configMapMapping(*cm), nil
	}
}

func (w *watcher) Close() error {
	close(w.closeChan)
	return nil
}
