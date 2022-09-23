package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
var (
	wsSubscribeCallStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ws_subscribe_call_status",
			Help: "ws subscribe call status."},
		[]string{"status", "address"},
	)
)

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(wsSubscribeCallStatus)
}
