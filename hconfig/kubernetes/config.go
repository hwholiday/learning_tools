package kubernetes

import (
	"errors"
	"github.com/hwholiday/learning_tools/hconfig/hconf"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var _ hconf.DataSource = (*kubernetesConfig)(nil)

type kubernetesConfig struct {
	client  *kubernetes.Clientset
	options *options
}

func NewKubernetesConfig(cli *kubernetes.Clientset, opts ...Option) (hconf.DataSource, error) {
	if cli == nil {
		return nil, errors.New("etcd client is nil")
	}
	conf := &kubernetesConfig{
		client:  cli,
		options: NewOptions(opts...),
	}
	return conf, nil
}

func (c *kubernetesConfig) configMapMapping(cm v1.ConfigMap) []*hconf.Data {
	var data = make([]*hconf.Data, 0)
	for key, val := range cm.Data {
		for _, path := range c.options.paths {
			if key == path {
				data = append(data, &hconf.Data{
					Key: key,
					Val: []byte(val),
				})
			}
		}
	}
	return data
}
func (c *kubernetesConfig) Load() ([]*hconf.Data, error) {
	ack, err := c.client.CoreV1().ConfigMaps(c.options.namespace).List(c.options.ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var data = make([]*hconf.Data, 0)
	for _, v := range ack.Items {
		data = append(data, c.configMapMapping(v)...)
	}
	return data, nil
}

func (c *kubernetesConfig) Watch() (hconf.DataWatcher, error) {
	return newWatcher(c)
}
