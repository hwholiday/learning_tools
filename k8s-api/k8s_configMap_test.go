package main

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestConfigMap(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/app/.kube/config"))
	if err != nil {
		panic(err)
	}
	fmt.Println(client == nil)
	data, err := json.Marshal(&map[string]interface{}{
		"name": 123,
		"test": "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
	ack, err := client.CoreV1().ConfigMaps("im").Create(context.Background(), &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "im-test-conf",
			Labels: map[string]string{
				"app": "test",
			},
		},
		Data: map[string]string{
			"gs-123-hw": string(data),
		},
	}, metav1.CreateOptions{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", ack)

}

func TestWatch(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/jk/.kube/config"))
	if err != nil {
		panic(err)
	}

	w, err := client.CoreV1().ConfigMaps("im").Watch(context.Background(), metav1.ListOptions{
		LabelSelector: "app=test",
		//FieldSelector: "metadata.name=im-test-conf",//metadata.name=im-test-conf
	})
	if err != nil {
		panic(err)
	}
	defer w.Stop()
	for {
		ch := <-w.ResultChan()
		cm, ok := ch.Object.(*v1.ConfigMap)
		if ok {
			fmt.Println("data", cm.Name, cm.Data)
		}
	}

}
