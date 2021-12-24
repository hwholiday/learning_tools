package main

import (
	"context"
	"encoding/json"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"testing"
	"time"
)

func TestK8s(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/app/conf/kube_config/local_kube.yaml"))
	if err != nil {
		panic(err)
	}

	informerFactory := informers.NewSharedInformerFactory(client, time.Minute)
	informer := informerFactory.Core().V1().Pods().Informer() //缓存
	podLister := informerFactory.Core().V1().Pods().Lister()
	var stopCh = make(chan struct{})
	var podCh = make(chan struct{})
	informerFactory.Start(stopCh)
	go func() {
		for {
			select {
			case <-podCh:
				pods, err := podLister.List(labels.SelectorFromSet(map[string]string{
					"app.kubernetes.io/instance": "rke2-ingress-nginx",
				}))
				if err != nil {
					panic(err)
				}
				fmt.Println("pods", len(pods))
				for _, v := range pods {
					if v.Status.Phase != corev1.PodRunning {
						continue
					}
					//data, err := json.Marshal(v)
					//if err != nil {
					//	panic(err)
					//}
					//fmt.Println("data", string(data))
					//t.Logf("服务地址 v.Status.PodIP %+v", v.Status.PodIP)
					for _, container := range v.Spec.Containers {
						for _, cp := range container.Ports {
							//t.Logf("cp %+v", cp.String())
							t.Logf("服务地址 PodIP %+v PORT %+v", v.Status.PodIP, cp.ContainerPort)
						}
					}
				}
			}
		}
	}()

	go func() {
		informer.AddEventHandler(cache.FilteringResourceEventHandler{
			FilterFunc: func(obj interface{}) bool {
				select {
				case <-stopCh:
					return false
				default:
					pod := obj.(*v1.Pod)
					val := pod.GetLabels()["app.kubernetes.io/instance"]
					return val == "rke2-ingress-nginx"
				}
			},
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {
					fmt.Println("AddFunc")
					//fmt.Println("AddFunc", obj)
					podCh <- struct{}{}
				},
				UpdateFunc: func(oldObj, newObj interface{}) {
					podCh <- struct{}{}
					fmt.Println("UpdateFunc")
					//fmt.Println("UpdateFunc", oldObj, newObj)
				},
				DeleteFunc: func(obj interface{}) {
					podCh <- struct{}{}
					fmt.Println("DeleteFunc")
					//fmt.Println("DeleteFunc", obj)
				},
			},
		})
	}()
	<-stopCh
	return
}

func TestSet(t *testing.T) {
	client, err := NewK8sClientset(KubeConfigPath("/home/app/conf/kube_config/local_kube.yaml"))
	if err != nil {
		panic(err)
	}
	patchBytes, err := json.Marshal(map[string]interface{}{
		"metadata": metav1.ObjectMeta{
			Labels: map[string]string{
				"labels-test": "labels-test2",
			},
			Annotations: map[string]string{
				"other-test": "other-test",
				"other-json": `{"a": "a", "b": 1}`,
			},
		},
	})
	res, err := client.CoreV1().Pods("kube-system").Patch(context.Background(), "rke2-ingress-nginx-controller-lvptn", types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("res %+v", res)

}
