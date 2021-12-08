package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// kubectl describe svc kubernetes

type Option func(*options)

type options struct {
	KubeConfigPath string
	MasterUrl      string
}

func KubeConfigPath(kubeConfigPath string) Option {
	return func(o *options) {
		o.KubeConfigPath = kubeConfigPath
	}
}

func MasterUrl(masterUrl string) Option {
	return func(o *options) {
		o.MasterUrl = masterUrl
	}
}

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
	if op.KubeConfigPath != "" {
		if config, err = clientcmd.BuildConfigFromFlags(op.MasterUrl, op.KubeConfigPath); err != nil {
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
