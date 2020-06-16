package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// service is online if set to 1, offline if 0
	databaseConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "database_connections",
			Help:      "If service is online",
		}, nil,
	)
	databaseInUse = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "database_connections_in_use",
			Help:      "If service is online",
		}, nil,
	)
	databaseIdle = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "database_connections_idle",
			Help:      "If service is online",
		}, nil,
	)
)

func Database(method string, value float64) {
	switch method {
	case "connections":
		databaseConnections.WithLabelValues().Set(value)
	case "in_use":
		databaseInUse.WithLabelValues().Set(value)
	case "idle":
		databaseIdle.WithLabelValues().Set(value)
	}
}
