package configs

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/checkins"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/groups"
	"github.com/hunterlong/statping/types/hits"
	"github.com/hunterlong/statping/types/incidents"
	"github.com/hunterlong/statping/types/integrations"
	"github.com/hunterlong/statping/types/messages"
	"github.com/hunterlong/statping/types/notifications"
	"github.com/hunterlong/statping/types/services"
	"github.com/hunterlong/statping/types/users"
	"github.com/hunterlong/statping/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type SamplerFunc func()

type Sampler interface {
	Samples() []database.DbObject
}

func TriggerSamples() error {
	return createSamples(
		core.Samples,
		users.Samples,
		messages.Samples,
		services.Samples,
		checkins.Samples,
		checkins.SamplesChkHits,
		failures.Samples,
		groups.Samples,
		hits.Samples,
		incidents.Samples,
		incidents.SamplesUpdates,
	)
}

func createSamples(sm ...SamplerFunc) error {
	for _, v := range sm {
		v()
	}
	return nil
}

func (d *DbConfig) Connect() error {

	return nil
}

func (d *DbConfig) Create() error {

	return nil
}

// Migrate function
func (d *DbConfig) Update() error {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		return err
	}
	defer config.Close()

	data, err := yaml.Marshal(d)
	if err != nil {
		log.Errorln(err)
		return err
	}
	config.WriteString(string(data))
	return nil
}

// Save will initially create the config.yml file
func (d *DbConfig) Delete() error {
	return os.Remove(d.filename)
}

// DropDatabase will DROP each table Statping created
func (d *DbConfig) DropDatabase() error {
	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}, &integrations.Integration{}}
	log.Infoln("Dropping Database Tables...")
	for _, t := range DbModels {
		if err := database.DB().DropTableIfExists(t); err != nil {
			return err.Error()
		}
		log.Infof("Dropped table: %T\n", t)
	}
	return nil
}

// CreateDatabase will CREATE TABLES for each of the Statping elements
func CreateDatabase() error {
	var err error

	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}, &integrations.Integration{}}

	log.Infoln("Creating Database Tables...")
	for _, table := range DbModels {
		if err := database.DB().CreateTable(table); err.Error() != nil {
			return err.Error()
		} else {
			log.Infof("Database table: '%T' was created", table)
		}
	}
	if err := database.DB().Table("core").CreateTable(&core.Core{}); err.Error() != nil {
		return err.Error()
	} else {
		log.Infof("Database table: '%s' was created", "core")
	}
	log.Infoln("Statping Database Created")

	return err
}
