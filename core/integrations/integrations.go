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
	"errors"
	"fmt"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
)

var (
	Integrations []types.Integrator
	log          = utils.Log.WithField("type", "integration")
	db           types.Database
)

//func init() {
//	Integrations = append(Integrations,
//		CsvIntegrator,
//		DockerIntegrator,
//		TraefikIntegrator,
//	)
//}

// integrationsDb returns the 'integrations' database column
func integrationsDb() types.Database {
	return db.Model(&types.Integration{})
}

// SetDB is called by core to inject the database for a integrator to use
func SetDB(d types.Database) {
	db = d
}

func Value(intg types.Integrator, fieldName string) interface{} {
	for _, v := range intg.Get().Fields {
		if fieldName == v.Name {
			return v.Value
		}
	}
	return nil
}

func Update(integrator *types.Integration) error {
	fields := FieldsToJson(integrator)
	fmt.Println(fields)
	set := db.Model(&types.Integration{}).Where("name = ?", integrator.Name)
	set.Set("enabled", integrator.Enabled)
	set.Set("fields", fields)
	return set.Error()
}

func FieldsToJson(integrator *types.Integration) string {
	jsonData := make(map[string]interface{})
	for _, v := range integrator.Fields {
		jsonData[v.Name] = v.Value
	}
	data, _ := json.Marshal(jsonData)
	return string(data)
}

func JsonToFields(intg types.Integrator, input string) []*types.IntegrationField {
	integrator := intg.Get()
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(input), &jsonData)

	for _, v := range integrator.Fields {
		v.Value = jsonData[v.Name]
	}
	return integrator.Fields
}

func SetFields(intg types.Integrator, data map[string][]string) (*types.Integration, error) {
	i := intg.Get()
	for _, v := range i.Fields {
		if data[v.Name] != nil {
			v.Value = data[v.Name][0]
		}
	}
	return i, nil
}

func Find(name string) (types.Integrator, error) {
	for _, i := range Integrations {
		obj := i.Get()
		if obj.ShortName == name {
			return i, nil
		}
	}
	return nil, errors.New(name + " not found")
}

// db will return the notifier database column/record
func integratorDb(n *types.Integration) types.Database {
	return db.Model(&types.Integration{}).Where("name = ?", n.Name).Find(n)
}

// isInDatabase returns true if the integration has already been installed
func isInDatabase(i types.Integrator) bool {
	inDb := integratorDb(i.Get()).RecordNotFound()
	return !inDb
}

// SelectIntegration returns the Notification struct from the database
func SelectIntegration(i types.Integrator) (*types.Integration, error) {
	integration := i.Get()
	err := db.Model(&types.Integration{}).Where("name = ?", integration.Name).Scan(&integration)
	return integration, err.Error()
}

// AddIntegrations accept a Integrator interface to be added into the array
func AddIntegrations(integrations ...types.Integrator) error {
	for _, i := range integrations {
		if utils.IsType(i, new(types.Integrator)) {
			Integrations = append(Integrations, i)
			err := install(i)
			if err != nil {
				return err
			}
		} else {
			return errors.New("notifier does not have the required methods")
		}
	}
	return nil
}

// install will check the database for the notification, if its not inserted it will insert a new record for it
func install(i types.Integrator) error {
	inDb := isInDatabase(i)
	log.WithField("installed", inDb).
		WithFields(utils.ToFields(i)).
		Debugln(fmt.Sprintf("Checking if integrator '%v' is installed: %v", i.Get().Name, inDb))
	if !inDb {
		_, err := insertDatabase(i)
		if err != nil {
			log.Errorln(err)
			return err
		}
	}
	return nil
}

// insertDatabase will create a new record into the database for the integrator
func insertDatabase(i types.Integrator) (string, error) {
	integrator := i.Get()
	query := db.Create(integrator)
	if query.Error() != nil {
		return "", query.Error()
	}
	return integrator.Name, query.Error()
}
