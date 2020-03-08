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

package integrators

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/integrations"
	"github.com/hunterlong/statping/utils"
)

var (
	Integrations []integrations.Integrator
	log          = utils.Log.WithField("type", "integration")
	db           database.Database
)

//func init() {
//	Integrations = append(Integrations,
//		CsvIntegrator,
//		DockerIntegrator,
//		TraefikIntegrator,
//	)
//}

func init() {
	AddIntegrations(
		CsvIntegrator,
		TraefikIntegrator,
		DockerIntegrator,
	)
}

// integrationsDb returns the 'integrations' database column
func integrationsDb() database.Database {
	return db.Model(&integrations.Integration{})
}

// SetDB is called by core to inject the database for a integrator to use
func SetDB(d database.Database) {
	db = d
}

func Value(intg integrations.Integrator, fieldName string) interface{} {
	for _, v := range intg.Get().Fields {
		if fieldName == v.Name {
			return v.Value
		}
	}
	return nil
}

func Update(integrator *integrations.Integration) error {
	fields := FieldsToJson(integrator)
	fmt.Println(fields)
	set := db.Model(&integrations.Integration{}).Where("name = ?", integrator.Name)
	set.Set("enabled", integrator.Enabled)
	set.Set("fields", fields)
	return set.Error()
}

func FieldsToJson(integrator *integrations.Integration) string {
	jsonData := make(map[string]interface{})
	for _, v := range integrator.Fields {
		jsonData[v.Name] = v.Value
	}
	data, _ := json.Marshal(jsonData)
	return string(data)
}

func JsonToFields(intg integrations.Integrator, input string) []*integrations.IntegrationField {
	integrator := intg.Get()
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(input), &jsonData)

	for _, v := range integrator.Fields {
		v.Value = jsonData[v.Name]
	}
	return integrator.Fields
}

func SetFields(intg integrations.Integrator, data map[string][]string) (*integrations.Integration, error) {
	i := intg.Get()
	for _, v := range i.Fields {
		if data[v.Name] != nil {
			v.Value = data[v.Name][0]
		}
	}
	return i, nil
}

func Find(name string) (integrations.Integrator, error) {
	for _, i := range Integrations {
		obj := i.Get()
		if obj.ShortName == name {
			return i, nil
		}
	}
	return nil, errors.New(name + " not found")
}

// db will return the notifier database column/record
func integratorDb(n *integrations.Integration) database.Database {
	return db.Model(&integrations.Integration{}).Where("name = ?", n.Name).Find(n)
}

// isInDatabase returns true if the integration has already been installed
func isInDatabase(i integrations.Integrator) bool {
	inDb := integratorDb(i.Get()).RecordNotFound()
	return !inDb
}

// SelectIntegration returns the Notification struct from the database
func SelectIntegration(i integrations.Integrator) (*integrations.Integration, error) {
	integration := i.Get()
	err := db.Model(&integrations.Integration{}).Where("name = ?", integration.Name).Scan(&integration)
	return integration, err.Error()
}

// AddIntegrations accept a Integrator interface to be added into the array
func AddIntegrations(inte ...integrations.Integrator) error {
	for _, i := range inte {
		if utils.IsType(i, new(integrations.Integrator)) {
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
func install(i integrations.Integrator) error {
	_, err := insertDatabase(i)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return err
}

// insertDatabase will create a new record into the database for the integrator
func insertDatabase(i integrations.Integrator) (string, error) {
	integrator := i.Get()
	query := db.FirstOrCreate(integrator)
	if query.Error() != nil {
		return "", query.Error()
	}
	return integrator.Name, query.Error()
}
