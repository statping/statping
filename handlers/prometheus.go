// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"fmt"
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
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

	secondsOnline := time.Now().Sub(utils.StartTime).Seconds()
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

	for _, n := range notifiers.All() {
		notif := n.Select()
		PrometheusComment(fmt.Sprintf("Notifier %s:", notif.Method))
		enabled := 0
		if notif.Enabled.Bool {
			enabled = 1
		}
		PrometheusExportKey("notifier_enabled", notif.Id, notif.Method, enabled)
		PrometheusExportKey("notifier_on_success", notif.Id, notif.Method, notif.Hits.OnSuccess)
		PrometheusExportKey("notifier_on_failure", notif.Id, notif.Method, notif.Hits.OnFailure)
		PrometheusExportKey("notifier_on_user_new", notif.Id, notif.Method, notif.Hits.OnNewUser)
		PrometheusExportKey("notifier_on_user_update", notif.Id, notif.Method, notif.Hits.OnUpdatedUser)
		PrometheusExportKey("notifier_on_user_delete", notif.Id, notif.Method, notif.Hits.OnDeletedUser)
		PrometheusExportKey("notifier_on_service_new", notif.Id, notif.Method, notif.Hits.OnNewService)
		PrometheusExportKey("notifier_on_service_update", notif.Id, notif.Method, notif.Hits.OnUpdatedService)
		PrometheusExportKey("notifier_on_service_delete", notif.Id, notif.Method, notif.Hits.OnDeletedService)
		PrometheusExportKey("notifier_on_notifier_new", notif.Id, notif.Method, notif.Hits.OnNewNotifier)
		PrometheusExportKey("notifier_on_notifier_update", notif.Id, notif.Method, notif.Hits.OnUpdatedNotifier)
		PrometheusExportKey("notifier_on_notifier_save", notif.Id, notif.Method, notif.Hits.OnSave)
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
