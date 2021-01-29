//gateway 为了简单我们使用以下协议
```base
  {
    "code": 1,//协议号
    "data": "请求内容",
  }
```
#### code = 1 请求 logic version
```base
{"code": 1,"data": ""}
```
#### code = 2 请求 ReqName
 ```base
{"code": 2,"data": "istio"}
```
#### code = 3 请求 测试ws服务
 ```base
{"code": 3,"data": "istio"}
```
#### 本地测试  ws://127.0.0.1:8888/connect
#### istio测试 ws://172.13.3.131:80/connect  （172.13.3.131为你的EXTERNAL-IP）
#### [在线测试WS地址](http://www.easyswoole.com/wstool.html)