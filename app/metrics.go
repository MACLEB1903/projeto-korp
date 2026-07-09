package main

import "github.com/prometheus/client_golang/prometheus"

var httpRequestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "projeto_korp_requests_total",
		Help: "Total number of requests received by the /projeto-korp endpoint.",
	},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}