// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/services"
	"net/http"
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

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	metrics := []string{}
	allFails := failures.All()
	system := fmt.Sprintf("statping_total_failures %v\n", allFails)
	system += fmt.Sprintf("statping_total_services %v", len(services.All()))
	metrics = append(metrics, system)
	for _, ser := range services.All() {
		online := 1
		if !ser.Online {
			online = 0
		}
		met := fmt.Sprintf("statping_service_failures{id=\"%v\" name=\"%v\"} %v\n", ser.Id, ser.Name, ser.AllFailures().Count())
		met += fmt.Sprintf("statping_service_latency{id=\"%v\" name=\"%v\"} %0.0f\n", ser.Id, ser.Name, (ser.Latency * 100))
		met += fmt.Sprintf("statping_service_online{id=\"%v\" name=\"%v\"} %v\n", ser.Id, ser.Name, online)
		met += fmt.Sprintf("statping_service_status_code{id=\"%v\" name=\"%v\"} %v\n", ser.Id, ser.Name, ser.LastStatusCode)
		met += fmt.Sprintf("statping_service_response_length{id=\"%v\" name=\"%v\"} %v", ser.Id, ser.Name, len([]byte(ser.LastResponse)))
		metrics = append(metrics, met)
	}
	output := strings.Join(metrics, "\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
