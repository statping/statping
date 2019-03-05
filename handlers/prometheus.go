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
	"github.com/hunterlong/statping/core"
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
	system := fmt.Sprintf("statping_total_failures %v\n", core.CountFailures())
	system += fmt.Sprintf("statping_total_services %v", len(core.CoreApp.Services))
	metrics = append(metrics, system)
	for _, ser := range core.CoreApp.Services {
		v := ser.Select()
		online := 1
		if !v.Online {
			online = 0
		}
		met := fmt.Sprintf("statping_service_failures{id=\"%v\" name=\"%v\"} %v\n", v.Id, v.Name, len(v.Failures))
		met += fmt.Sprintf("statping_service_latency{id=\"%v\" name=\"%v\"} %0.0f\n", v.Id, v.Name, (v.Latency * 100))
		met += fmt.Sprintf("statping_service_online{id=\"%v\" name=\"%v\"} %v\n", v.Id, v.Name, online)
		met += fmt.Sprintf("statping_service_status_code{id=\"%v\" name=\"%v\"} %v\n", v.Id, v.Name, v.LastStatusCode)
		met += fmt.Sprintf("statping_service_response_length{id=\"%v\" name=\"%v\"} %v", v.Id, v.Name, len([]byte(v.LastResponse)))
		metrics = append(metrics, met)
	}
	output := strings.Join(metrics, "\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
