package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestK8s(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/jk/.kube/config"))
	if err != nil {
		panic(err)
	}
	pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, v := range pods.Items {
		t.Logf("pods %+v", v)
	}
}
