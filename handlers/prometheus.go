package handlers

import (
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

//
//   Use your Statping Secret API Key for Authentication
//
//		scrape_configs:
//		  - job_name: 'statping'
//			bearer_token: MY API SECRET HERE
//			static_configs:
//			  - targets: ['statping:8080']
//

var (
	prefix       string
	promValues   []string
	httpRequests int64
	httpBytesIn  int64
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func hex2int(hexStr string) uint64 {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	promValues = []string{}
	prefix = utils.Getenv("PREFIX", "").(string)
	if prefix != "" {
		prefix = prefix + "_"
	}

	secondsOnline := utils.Now().Sub(utils.StartTime).Seconds()
	allFails := failures.All()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	httpMetrics := utils.GetHttpMetrics()

	promValues = append(promValues, "# Statping Prometheus Exporter")

	PrometheusComment("Statping Totals")
	PrometheusKeyValue("total_failures", len(allFails))
	PrometheusKeyValue("total_services", len(services.Services()))
	PrometheusKeyValue("seconds_online", secondsOnline)

	if secondsOnline < 5 {
		return
	}

	for _, ser := range services.AllInOrder() {
		online := 1
		if !ser.Online {
			online = 0
		}
		id := ser.Id
		name := ser.Name

		PrometheusComment(fmt.Sprintf("Service #%d '%s':", ser.Id, ser.Name))
		PrometheusExportKey("service_failures", id, name, ser.AllFailures().Count())
		PrometheusExportKey("service_latency", id, name, ser.Latency)
		PrometheusExportKey("service_online", id, name, online)
		PrometheusExportKey("service_status_code", id, name, ser.LastStatusCode)
		PrometheusExportKey("service_response_length", id, name, len([]byte(ser.LastResponse)))
		PrometheusExportKey("service_ping_time", id, name, ser.PingTime)
		PrometheusExportKey("service_last_latency", id, name, ser.LastLatency)
		PrometheusExportKey("service_last_lookup", id, name, ser.LastLookupTime)
		PrometheusExportKey("service_last_check", id, name, utils.Now().Sub(ser.LastCheck).Milliseconds())
		//PrometheusExportKey("service_online_seconds", id, name, ser.SecondsOnline)
		//PrometheusExportKey("service_offline_seconds", id, name, ser.SecondsOffline)

	}

	for _, n := range services.AllNotifiers() {
		notif := n.Select()
		if notif.Enabled.Bool {
			PrometheusComment(fmt.Sprintf("Notifier %s:", notif.Method))
			PrometheusNoIDExportKey("notifier_on_success", notif.Method, notif.Hits.OnSuccess)
			PrometheusNoIDExportKey("notifier_on_failure", notif.Method, notif.Hits.OnFailure)
			PrometheusNoIDExportKey("notifier_on_user_new", notif.Method, notif.Hits.OnNewUser)
			PrometheusNoIDExportKey("notifier_on_user_update", notif.Method, notif.Hits.OnUpdatedUser)
			PrometheusNoIDExportKey("notifier_on_user_delete", notif.Method, notif.Hits.OnDeletedUser)
			PrometheusNoIDExportKey("notifier_on_service_new", notif.Method, notif.Hits.OnNewService)
			PrometheusNoIDExportKey("notifier_on_service_update", notif.Method, notif.Hits.OnUpdatedService)
			PrometheusNoIDExportKey("notifier_on_service_delete", notif.Method, notif.Hits.OnDeletedService)
			PrometheusNoIDExportKey("notifier_on_notifier_new", notif.Method, notif.Hits.OnNewNotifier)
			PrometheusNoIDExportKey("notifier_on_notifier_update", notif.Method, notif.Hits.OnUpdatedNotifier)
			PrometheusNoIDExportKey("notifier_on_notifier_save", notif.Method, notif.Hits.OnSave)
		}
	}

	PrometheusComment("HTTP Metrics")
	PrometheusKeyValue("http_errors", httpMetrics.Errors)
	PrometheusKeyValue("http_requests", httpMetrics.Requests)
	PrometheusKeyValue("http_bytes", httpMetrics.Bytes)
	PrometheusKeyValue("http_request_milliseconds", httpMetrics.Milliseconds)

	// https://golang.org/pkg/runtime/#MemStats
	PrometheusComment("Golang Metrics")
	PrometheusKeyValue("go_heap_allocated", m.Alloc)
	PrometheusKeyValue("go_total_allocated", m.TotalAlloc)
	PrometheusKeyValue("go_heap_in_use", m.HeapInuse)
	PrometheusKeyValue("go_heap_objects", m.HeapObjects)
	PrometheusKeyValue("go_heap_idle", m.HeapIdle)
	PrometheusKeyValue("go_heap_released", m.HeapReleased)
	PrometheusKeyValue("go_heap_frees", m.Frees)
	PrometheusKeyValue("go_lookups", m.Lookups)
	PrometheusKeyValue("go_system", m.Sys)
	PrometheusKeyValue("go_number_gc", m.NumGC)
	PrometheusKeyValue("go_number_gc_forced", m.NumForcedGC)
	PrometheusKeyValue("go_goroutines", runtime.NumGoroutine())

	output := strings.Join(promValues, "\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func PrometheusKeyValue(keyName string, value interface{}) {
	val := promValue(value)
	prom := fmt.Sprintf("%sstatping_%s %s", prefix, keyName, val)
	promValues = append(promValues, prom)
}

func PrometheusExportKey(keyName string, id int64, name string, value interface{}) {
	val := promValue(value)
	prom := fmt.Sprintf("%sstatping_%s{id=\"%d\" name=\"%s\"} %s", prefix, keyName, id, name, val)
	promValues = append(promValues, prom)
}

func PrometheusNoIDExportKey(keyName string, name string, value interface{}) {
	val := promValue(value)
	prom := fmt.Sprintf("%sstatping_%s{name=\"%s\"} %s", prefix, keyName, name, val)
	promValues = append(promValues, prom)
}

func PrometheusComment(comment string) {
	prom := fmt.Sprintf("\n# %v", comment)
	promValues = append(promValues, prom)
}

func promValue(val interface{}) string {
	var newVal string
	switch v := val.(type) {
	case float64:
		newVal = fmt.Sprintf("%.4f", v)
	default:
		newVal = fmt.Sprintf("%v", v)
	}
	return newVal
}
