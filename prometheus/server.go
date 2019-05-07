package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
	"math/rand"
	"fmt"
)

var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count"},
	[]string{"endpoint"})

var orderNum = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "order_num",
		Help: "order num"})

var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(orderNum)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/howie", howie)
	fmt.Println("服务器启动192.168.2.28:8888")
	err := http.ListenAndServe("192.168.2.28:8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func howie(w http.ResponseWriter, r *http.Request) {
	httpRequestCount.WithLabelValues(r.URL.Path).Inc()
	start := time.Now()
	n := rand.Intn(100)
	if n >= 90 {
		orderNum.Dec()
		time.Sleep(100 * time.Millisecond)
	} else {
		orderNum.Inc()
		time.Sleep(50 * time.Millisecond)
	}

	elapsed := (float64)(time.Since(start) / time.Millisecond)
	httpRequestDuration.WithLabelValues(r.URL.Path).Observe(elapsed)
	w.Write([]byte("ok"))
}
