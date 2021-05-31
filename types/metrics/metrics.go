package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	utilsHttpRequestDur = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "statping",
			Name:      "http_requests_duration",
			Help:      "Duration for h",
		},
		[]string{"url", "method"},
	)

	utilsHttpRequestBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "statping",
			Name:      "http_response_bytes",
			Help:      "Response in bytes for a HTTP services",
		},
		[]string{"url", "method"},
	)

	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "statping",
			Name:      "http_duration_seconds",
			Help:      "Duration of HTTP requests from the utils package",
		}, []string{"path"})
)

func InitMetrics() {
	prometheus.MustRegister(
		serviceOnline,
		serviceFailures,
		serviceSuccess,
		serviceStatusCode,
		serviceDuration,
		utilsHttpRequestDur,
		utilsHttpRequestBytes,
		httpDuration,
		databaseStats,
		queryStats,
	)
}

func Histo(method string, value float64, labels ...interface{}) {
	switch method {
	case "duration":
		utilsHttpRequestDur.WithLabelValues(convert(labels)...).Observe(value)
	case "bytes":
		utilsHttpRequestBytes.WithLabelValues(convert(labels)...).Observe(value)
	}
}

func Timer(labels ...interface{}) prometheus.Observer {
	return httpDuration.WithLabelValues(convert(labels)...)
}

func ServiceTimer(labels ...interface{}) prometheus.Observer {
	return serviceDuration.WithLabelValues(convert(labels)...)
}

func Gauge(method string, value float64, labels ...interface{}) {
	switch method {
	case "status_code":
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
