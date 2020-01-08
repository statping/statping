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

package integrations

import (
	"encoding/json"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/url"
	"time"
)

type traefikIntegration struct {
	*types.Integration
}

var traefikIntegrator = &traefikIntegration{&types.Integration{
	ShortName:   "traefik",
	Name:        "Traefik",
	Icon:        "<i class=\"fas fa-network-wired\"></i>",
	Description: ``,
	Fields: []*types.IntegrationField{
		{
			Name:        "endpoint",
			Description: "The URL for the traefik API Endpoint",
			Type:        "text",
			Value:       "http://localhost:8080",
		},
		{
			Name:        "username",
			Description: "Username for HTTP Basic Authentication",
			Type:        "text",
		},
		{
			Name:        "password",
			Description: "Password for HTTP Basic Authentication",
			Type:        "password",
		},
	},
}}

func (t *traefikIntegration) Get() *types.Integration {
	return t.Integration
}

func (t *traefikIntegration) List() ([]*types.Service, error) {
	var err error
	var services []*types.Service

	endpoint := Value(t, "endpoint").(string)

	httpServices, err := fetchMethod(endpoint, "http")
	if err != nil {
		return nil, err
	}
	services = append(services, httpServices...)

	tcpServices, err := fetchMethod(endpoint, "tcp")
	if err != nil {
		return nil, err
	}
	services = append(services, tcpServices...)

	return services, err
}

func fetchMethod(endpoint, method string) ([]*types.Service, error) {
	var traefikServices []traefikService
	var services []*types.Service
	d, _, err := utils.HttpRequest(endpoint+"/api/"+method+"/services", "GET", nil, []string{}, nil, 10*time.Second, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(d, &traefikServices); err != nil {
		return nil, err
	}
	for _, s := range traefikServices {
		log.Infoln(s)

		for _, l := range s.LoadBalancer.Servers {

			url, err := url.Parse(l.URL)
			if err != nil {
				return nil, err
			}

			service := &types.Service{
				Name:     s.Name,
				Domain:   url.Hostname(),
				Port:     int(utils.ToInt(url.Port())),
				Type:     method,
				Interval: 60,
				Timeout:  2,
			}
			services = append(services, service)

		}
	}
	return services, err
}

type traefikService struct {
	Status       string   `json:"status"`
	UsedBy       []string `json:"usedBy"`
	Name         string   `json:"name"`
	Provider     string   `json:"provider"`
	LoadBalancer struct {
		Servers []struct {
			URL string `json:"url"`
		} `json:"servers"`
		PassHostHeader bool `json:"passHostHeader"`
	} `json:"loadBalancer,omitempty"`
}
