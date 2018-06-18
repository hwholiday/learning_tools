package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	http.ListenAndServe("8080", nil)
}

func howie(w http.ResponseWriter,r *http.Request)  {
}
