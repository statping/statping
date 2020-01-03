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
	"errors"
	"fmt"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const requiredSize = 17

type csvIntegration struct {
	*types.Integration
}

var csvIntegrator = &csvIntegration{&types.Integration{
	ShortName:   "csv",
	Name:        "CSV File",
	Description: "Import multiple services from a CSV file",
	Fields: []*types.IntegrationField{
		{
			Name:     "input",
			Type:     "file",
			MimeType: "application/csv",
		},
	},
}}

var csvData [][]string

func (t *csvIntegration) Get() *types.Integration {
	return t.Integration
}

func (t *csvIntegration) List() ([]*types.Service, error) {
	path := csvIntegrator.Fields[0].Value.(string)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(strings.TrimSuffix(string(data), "\n"), "\n") {
		col := strings.Split(line, ",")
		csvData = append(csvData, col)
	}

	var services []*types.Service
	for _, v := range csvData {
		s, err := commaToService(v)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

// commaToService will convert a CSV comma delimited string slice to a Service type
// this function is used for the bulk import services feature
func commaToService(s []string) (*types.Service, error) {
	if len(s) != requiredSize {
		err := fmt.Errorf("file has %v columns of data, not the expected amount of %v columns for a service", len(s), requiredSize)
		return nil, err
	}

	interval, err := time.ParseDuration(s[4])
	if err != nil {
		return nil, errors.New("could not parse internal duration: " + s[4])
	}

	timeout, err := time.ParseDuration(s[9])
	if err != nil {
		return nil, errors.New("could not parse timeout duration: " + s[9])
	}

	allowNotifications, err := strconv.ParseBool(s[11])
	if err != nil {
		return nil, errors.New("could not parse allow notifications boolean: " + s[11])
	}

	public, err := strconv.ParseBool(s[12])
	if err != nil {
		return nil, errors.New("could not parse public boolean: " + s[12])
	}

	verifySsl, err := strconv.ParseBool(s[16])
	if err != nil {
		return nil, errors.New("could not parse verifiy SSL boolean: " + s[16])
	}

	newService := &types.Service{
		Name:               s[0],
		Domain:             s[1],
		Expected:           types.NewNullString(s[2]),
		ExpectedStatus:     int(utils.ToInt(s[3])),
		Interval:           int(utils.ToInt(interval.Seconds())),
		Type:               s[5],
		Method:             s[6],
		PostData:           types.NewNullString(s[7]),
		Port:               int(utils.ToInt(s[8])),
		Timeout:            int(utils.ToInt(timeout.Seconds())),
		AllowNotifications: types.NewNullBool(allowNotifications),
		Public:             types.NewNullBool(public),
		GroupId:            int(utils.ToInt(s[13])),
		Headers:            types.NewNullString(s[14]),
		Permalink:          types.NewNullString(s[15]),
		VerifySSL:          types.NewNullBool(verifySsl),
	}

	return newService, nil

}
