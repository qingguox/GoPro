package main

import (
	. "fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/actuator/prometheus", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}

// // prometheus 包是 prometheus/client_golang 的核心包
// go get /prometheus/client_golang/prometheus
// // promauto 包提供 Prometheus 指标的基本数据类型
// go get /prometheus/client_golang/prometheus/promauto
// // promhttp 包提供了 HTTP 服务端和客户端相关工具
// go get /prometheus/client_golang/prometheus/promhttp

// scrape_configs:
//  - job_name: "node_exporter"
//     static_configs:
//     - targets: ["localhost:9100"]
//  - job_name: "goPro"
//     metrics_path: '/actuator/prometheus'
//     static_configs:
//     - targets: ["localhost:8090"]  // 这里需要真实的地址 用 下面的mac进行查询即可

// mac 精准查询ip地址
// ifconfig en0 | grep inet| awk '{print $2}'

//   docker run  -d --name prometheus -p 9090:9090 -v /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

//   docker run  -d --name node_exporter -p 9100:9100  prom/node-exporter
