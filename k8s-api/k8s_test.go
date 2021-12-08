package main

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"testing"
	"time"
)

func TestK8s(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/jk/.kube/config"))
	if err != nil {
		panic(err)
	}

	informerFactory := informers.NewSharedInformerFactory(client, time.Minute*10)
	informerFactory.Core().V1().Pods().Informer()
	podLister := informerFactory.Core().V1().Pods().Lister()
	var stopCh = make(chan struct{})
	informerFactory.Start(stopCh)
	//pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
	//	LabelSelector: labels.FormatLabels(map[string]string{
	//		"component": "etcd",
	//	}),
	//})
	go func() {
		for {

			pods, err := podLister.List(labels.SelectorFromSet(map[string]string{
				//"component": "etcd",
			}))
			if err != nil {
				panic(err)
			}
			fmt.Println("pods", len(pods))
			for _, v := range pods {
				data, err := json.Marshal(v)
				if err != nil {
					panic(err)
				}
				fmt.Println("data", string(data))
				t.Logf("pods %+v", v.Status.PodIPs)
				for _, container := range v.Spec.Containers {
					for _, cp := range container.Ports {
						t.Logf("cp %+v", cp.String())
						port := cp.ContainerPort
						fmt.Println("port", port)
					}
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()

	<-stopCh
	return
}

func TestSet(t *testing.T) {
}
