package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"hash/crc32"
	"learning_tools/go-kit/v7/user_agent/pb"
	"learning_tools/go-kit/v7/user_agent/src"
	"learning_tools/go-kit/v7/utils"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var grpcAddr = flag.String("g", "127.0.0.1:8881", "grpcAddr")
var prometheusAddr = flag.String("p", "192.168.2.28:10001", "prometheus addr")

var quitChan = make(chan error, 1)

func main() {
	flag.Parse()
	var (
		etcdAddrs = []string{"127.0.0.1:2379"}
		serName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	utils.NewLoggerServer()
	//初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		utils.GetLogger().Error("[user_agent]  NewClient", zap.Error(err))
		return
	}
	Registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", serName, crc32.ChecksumIEEE([]byte(*grpcAddr))),
		Value: *grpcAddr,
	}, log.NewNopLogger())
	go func() {
		count := metricsprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "Number of requests",
		}, []string{"method"})

		histogram := metricsprometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "Request consumes time",
		}, []string{"method"})
		golangLimit := rate.NewLimiter(10, 1)
		server := src.NewService(utils.GetLogger(), count, histogram)
		endpoints := src.NewEndPointServer(server, golangLimit)
		grpcServer := src.NewGRPCServer(endpoints, utils.GetLogger())
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			utils.GetLogger().Warn("[user_agent] Listen", zap.Error(err))
			quitChan <- err
			return
		}
		Registar.Register()
		utils.GetLogger().Info("[user_agent] grpc run " + *grpcAddr)
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
		pb.RegisterUserServer(baseServer, grpcServer)
		quitChan <- baseServer.Serve(grpcListener)
	}()
	go func() {
		utils.GetLogger().Info("[user_agent] prometheus run " + *prometheusAddr)
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(*prometheusAddr, m)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()
	err = <-quitChan
	Registar.Deregister()
	utils.GetLogger().Info("[user_agent] quit", zap.Any("info", err))
}
