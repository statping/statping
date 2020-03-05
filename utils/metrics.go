package utils

import "time"

func init() {
	httpMetric = new(Metrics)
}

var (
	httpMetric *Metrics
	StartTime  = time.Now()
)

type Metrics struct {
	Requests     int64
	Errors       int64
	Bytes        int64
	Milliseconds int64
	OnlineTime   time.Time
}

func (h *Metrics) Reset() {
	httpMetric = new(Metrics)
}

func GetHttpMetrics() *Metrics {
	defer httpMetric.Reset()
	return httpMetric
}
