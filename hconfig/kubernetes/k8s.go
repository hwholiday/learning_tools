package kubernetes

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func NewK8sClientset(opts ...Option) (*kubernetes.Clientset, error) {
	op := options{}
	for _, o := range opts {
		o(&op)
	}
	var (
		config    *rest.Config
		err       error
		clientSet *kubernetes.Clientset
	)
	if op.kubeConfigPath != "" {
		if config, err = clientcmd.BuildConfigFromFlags(op.masterUrl, op.kubeConfigPath); err != nil {
			return nil, err
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			return nil, err
		}
	}
	if clientSet, err = kubernetes.NewForConfig(config); err != nil {
		return nil, err
	}
	return clientSet, err
}

const ServiceAccount = "/var/run/secrets/kubernetes.io/serviceaccount"

var namespace = LoadNamespace()

func LoadNamespace() string {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", ServiceAccount, "namespace"))
	if err != nil {
		return ""
	}
	return string(data)
}

// GetNamespace get k8s namespace
func GetNamespace() string {
	return namespace
}

// GetPodName get k8s pod name
func GetPodName() string {
	return os.Getenv("HOSTNAME")
}
