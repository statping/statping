package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	utilsHttpRequestDur = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_requests_duration",
			Help: "How many successful requests for a service",
		},
		[]string{"url", "method"},
	)

	utilsHttpRequestBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_bytes",
			Help: "How many successful requests for a service",
		},
		[]string{"url", "method"},
	)
)

func init() {
	prometheus.MustRegister(
		serviceOnline,
		serviceFailures,
		serviceSuccess,
		serviceStatusCode,
		serviceLatencyDuration,
		utilsHttpRequestDur,
		utilsHttpRequestBytes,
	)
}

func Histo(method string, value float64, labels ...interface{}) {
	switch method {
	case "latency":
		serviceLatencyDuration.WithLabelValues(convert(labels)...).Observe(value)
	case "duration":
		utilsHttpRequestDur.WithLabelValues(convert(labels)...).Observe(value)
	case "bytes":
		utilsHttpRequestBytes.WithLabelValues(convert(labels)...).Observe(value)
	}
}

func Gauge(method string, value float64, labels ...interface{}) {
	switch method {
	case "service":
		serviceStatusCode.WithLabelValues(convert(labels)...).Set(value)
	case "online":
		serviceOnline.WithLabelValues(convert(labels)...).Set(value)
	}
}

func Inc(method string, labels ...interface{}) {
	switch method {
	case "failure":
		serviceFailures.WithLabelValues(convert(labels)...).Inc()
	case "success":
		serviceSuccess.WithLabelValues(convert(labels)...).Inc()
	}
}

func Add(method string, value float64, labels ...interface{}) {
	switch method {
	case "failure":
		serviceFailures.WithLabelValues(convert(labels)...).Add(value)
	case "success":
		serviceSuccess.WithLabelValues(convert(labels)...).Add(value)
	}
}

func convert(vals []interface{}) []string {
	var out []string
	for _, v := range vals {
		out = append(out, fmt.Sprintf("%v", v))
	}
	return out
}
