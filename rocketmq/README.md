#### docker run --name rmqnamesrv -d -p 9876:9876 rocketmqinc/rocketmq:latest sh mqnamesrv
 


#### broker.conf
```base
terName = DefaultCluster
brokerName = broker-a
brokerId = 0
deleteWhen = 04
fileReservedTime = 48
brokerRole = ASYNC_MASTER
flushDiskType = ASYNC_FLUSH
brokerIP1 = 172.13.3.160
autoCreateTopicEnable=true
```


#### docker run --name rmqbroker -d -p 10911:10911 -p 10909:10909  --link rmqnamesrv:namesrv -e "NAMESRV_ADDR=namesrv:9876"  -v /root/docker_data/broker.conf:/opt/rocketmq/conf/broker.conf rocketmqinc/rocketmq:latest sh mqbroker -c  /opt/rocketmq/conf/broker.conf



#### docker run --name rmqconsole -d -p 8080:8080 --link rmqnamesrv:namesrv -e "JAVA_OPTS=-Drocketmq.namesrv.addr=namesrv:9876"  pangliang/rocketmq-console-ng