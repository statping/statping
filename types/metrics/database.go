package metrics

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// service is online if set to 1, offline if 0
	databaseStats = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "statping",
			Name:      "database",
			Help:      "If service is online",
		}, []string{"metric"},
	)

	queryStats = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "statping",
			Name:      "query",
			Help:      "If service is online",
		}, []string{"type", "method"},
	)
)

func Query(objType, method string) {
	queryStats.WithLabelValues(objType, method).Inc()
}

func CollectDatabase(stats sql.DBStats) {
	databaseStats.WithLabelValues("max_open_connections").Set(float64(stats.MaxOpenConnections))
	databaseStats.WithLabelValues("open_connections").Set(float64(stats.OpenConnections))
	databaseStats.WithLabelValues("in_use_connections").Set(float64(stats.InUse))
	databaseStats.WithLabelValues("idle_connections").Set(float64(stats.Idle))
	databaseStats.WithLabelValues("wait_count").Set(float64(stats.WaitCount))
	databaseStats.WithLabelValues("wait_duration_seconds").Set(stats.WaitDuration.Seconds())
	databaseStats.WithLabelValues("idle_connections_closed").Set(float64(stats.MaxIdleClosed))
	databaseStats.WithLabelValues("lifetime_connections_closed").Set(float64(stats.MaxLifetimeClosed))
}
