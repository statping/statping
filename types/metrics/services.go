package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// service is online if set to 1, offline if 0
	serviceOnline = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_online",
			Help: "If service is online",
		},
		[]string{"service"},
	)

	// service failures
	serviceFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_failures",
			Help: "How many failures occur for a service",
		},
		[]string{"service"},
	)

	// successful hits for a service
	serviceSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_success",
			Help: "How many successful requests for a service",
		},
		[]string{"service"},
	)

	// service check latency
	serviceLatencyDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "service_latency",
			Help: "How many successful requests for a service",
		},
		[]string{"service"},
	)

	// http status code for a service
	serviceStatusCode = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_status_code",
			Help: "HTTP Status code for a service",
		},
		[]string{"service"},
	)
)
