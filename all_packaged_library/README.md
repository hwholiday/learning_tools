# 实用工具库

### 基于etcd实现的服务注册，发现，负载均衡
```base
  	var op = SelectorOptions{
  		name: "svc.info",
  		config: clientv3.Config{
  			Endpoints:   []string{"http://localhost:2379/"},
  			DialTimeout: 5 * time.Second},
  	}
  	s, err := NewSelector(op)
  	if err != nil {
  		t.Error(err)
  		return
  	}
  	for {
  		val, err := s.Next()
  		if err != nil {
  			t.Error(err)
  			continue
  		}
  		fmt.Println(val)
  		time.Sleep(time.Second * 2)
  	}
```

### perf 　开启http pprof进行系统监控
```base
    perf.StartPprof([]string{"127.0.0.1:8077"})
```
### quit  优雅的退出go程序 
```base
    quit.QuitSignal(func() {
	    fmt.Println("退出程序")
    })
```
### zap 日志库相关
```base
    data := &Options{}
    	data.Development = true
    	initLogger(data)
    	for i := 0; i < 10; i++ {
    		time.Sleep(time.Second)
    		GetLogger().Debug(fmt.Sprint("debug log ", i), zap.Int("line", 47))
    		GetLogger().Info(fmt.Sprint("Info log ", i), zap.Any("level", "1231231231"))
    		GetLogger().Warn(fmt.Sprint("warn log ", i), zap.String("level", `{"a":"4","b":"5"}`))
    		GetLogger().Error(fmt.Sprint("err log ", i), zap.String("level", `{"a":"7","b":"8"}`))
    	}
```

### 华为推送
```base
    push := NewHuaweiPush("https://login.cloud.huawei.com/oauth2/v2/token", "100358845", "bee1d8f704b1bc278bea7f5427cb0f8a", "https://api.push.hicloud.com/pushsend.do", true)
    	var in ReqPush
    	in.DeviceTokenList = []string{"0862791036717594300002894200CN01"}
    	in.Ver = "1"
    	in.NspTs = "1545113076"
    	in.Payload = `{"hps":{"msg":{"type":1,"body":{"key":"value"}}}} `
    	for {
    		time.Sleep(time.Second * 3)
    		//判断ResPush中的code是不是等于80000000可以测试是否成功
    		//{"requestId":"154536200334112580200","msg":"Success","code":"80000000"}
    		fmt.Println(push.Push(&in))
    }
```

### 苹果pushkit推送
```base
   push,err:=InitPushKit("./131231P.p12","pwd",true)
   	if err!=nil{
   		fmt.Println(err)
   	}
   	fmt.Println(push.Push("123123",[]byte(`{"newsid":{"content":"test",},"badge":{"badge":"0"}}`)))
```


