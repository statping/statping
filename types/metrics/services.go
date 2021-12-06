package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// service is online if set to 1, offline if 0
	serviceOnline = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "service_online",
			Help:      "If service is online",
		},
		[]string{"service", "type"},
	)

	// service failures
	serviceFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "statping",
			Name:      "service_failures",
			Help:      "How many failures occur for a service",
		},
		[]string{"service"},
	)

	// successful hits for a service
	serviceSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "statping",
			Name:      "service_success",
			Help:      "How many successful requests for a service",
		},
		[]string{"service"},
	)

	// service check latency
	serviceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "statping",
			Name:      "service_duration",
			Help:      "Service request duration for a success response",
		},
		[]string{"service"},
	)

	// http status code for a service
	serviceStatusCode = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "service_status_code",
			Help:      "HTTP Status code for a service",
		},
		[]string{"service"},
	)
)
