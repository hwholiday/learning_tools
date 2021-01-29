## 安装 docker

- https://docs.docker.com/engine/install/

## 安装单节点 k8s 环境

```base
sudo docker run --privileged -d --restart=unless-stopped -p 88:80 -p 433:443 rancher/rancher
访问 https://宿主机IP:433
新建集群单结点要把 etcd ， control plane ，worker 都给勾选上
复制命令启动，等待就好了
```
## 软件架构
![](https://i.bmp.ovh/imgs/2021/01/218b88e77fa3018d.png)

## 准备镜像

- 网关：ws 服务
    - [在目录 gateway 下](https://github.com/hwholiday/learning_tools/tree/master/istio/gateway)
- 服务 V1:grpc 服务
    - [在目录 logic_v1 下](https://github.com/hwholiday/learning_tools/tree/master/istio/logic_v1)
- 服务 V2:grpc 服务
    - [在目录 logic_v2 下](https://github.com/hwholiday/learning_tools/tree/master/istio/logic_v2)
- 服务 V3:grpc 服务
    - [在目录 logic_v3 下](https://github.com/hwholiday/learning_tools/tree/master/istio/logic_v3)
    - 链接了 Reids 集群
- 关于k8s和istio的配置文件
  - [在目录 kube 下](https://github.com/hwholiday/learning_tools/tree/master/istio/kube)
      ```base
      ├── gateway.yaml      部署 gateway(Deployment[k8s]) 服务和对应的 Service(k8s)
      ├── logic.yaml        部署 logic(Deployment[k8s]) 服务和对应的 Service(k8s)
      ├── net-gateway.yaml  部署 Gateway(istio) VirtualService(istio) 接受外部请求转发到 gateway 服务
      ├── net-logic.yaml    部署 VirtualService(istio) DestinationRule(istio) 接受 gateway 请求转发到 logic 服务
      └── net-redis.yaml    部署 Service(k8s) Endpoints(k8s) 将外部服务转发到内部
      ```

## 安装 istio

- https://istio.io/ 建议看英文文档，中文的更新不及时可以对照着看

```base
istioctl install 

kubectl create namespace im

kubectl label namespace im istio-injection=enabled //在这个命名空间下的服务都会被自动注入 istio proxy

kubectl apply -f <(istioctl kube-inject -f 你的配置文件.yaml) -n 命令空间   //手动注入

```
- 查看是否有 EXTERNA-IP

```base
kubectl get service -n istio-system                                                                                                                                                     
NAME                   TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   LoadBalancer   10.43.50.213   <pending>     15021:31954/TCP,80:32725/TCP,443:32687/TCP,15012:30629/TCP,15443:32721/TCP   20m

#如果看到 EXTERNAL-IP 不为真实IP
# 手动指定边缘网络IP在 clusterIP 下输入
kubectl edit service/istio-ingressgateway -n istio-system

externalIPs:
 - 172.13.3.131

kubectl get service -n istio-system                                                                                                                                                     
NAME                   TYPE           CLUSTER-IP     EXTERNAL-IP    PORT(S)                                                                      AGE
istio-ingressgateway   LoadBalancer   10.43.50.213   172.13.3.131   15021:31954/TCP,80:32725/TCP,443:32687/TCP,15012:30629/TCP,15443:32721/TCP   26m
```

- 安装 kiali

```base
kubectl apply -f samples/addons
kubectl rollout status deployment/kiali -n istio-system
istioctl dashboard kiali
```  

- 安装 gateway 切换到learnin_tools/istio/kube目录下

```base
  kubectl apply -f gateway.yaml -n im
  kubectl apply -f net-gateway.yaml -n im
  
  //更新镜像
  //kubectl set image deploy gateway-v3 gateway=hwholiday/gateway:v3 -n im
  
  kubectl apply -f logic.yaml -n im
  kubectl apply -f net-logic.yaml -n im
  kubectl apply -f net-redis.yaml -n im
```
### [测试 gateway](https://github.com/hwholiday/learning_tools/blob/master/istio/gateway/README.md)

## 其他

- Service的DNS(host) 服务名.命名空间.svc.cluster.local

```base
gateway-v3-7d85c6577f-8xgwz 为pod name

kubectl describe pod/gateway-v3-7d85c6577f-8xgwz -n im

在 istio-proxy 标签里面能看到
```

- 查看 gateway 路由规则

```base
istioctl proxy-config routes istio-ingressgateway-865d46c7f5-2qgdh -n istio-system -o json
```

- 部署中的自动缩放
```base
kubectl autoscale deployment logic-v1 --cpu-percent=50 --min=1 --max=10 -n im
kubectl get hpa -n im
```

