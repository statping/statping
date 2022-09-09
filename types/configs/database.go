package configs

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/razorpay/statping/database"
	"github.com/razorpay/statping/types/checkins"
	"github.com/razorpay/statping/types/core"
	"github.com/razorpay/statping/types/failures"
	"github.com/razorpay/statping/types/groups"
	"github.com/razorpay/statping/types/hits"
	"github.com/razorpay/statping/types/incidents"
	"github.com/razorpay/statping/types/messages"
	"github.com/razorpay/statping/types/notifications"
	"github.com/razorpay/statping/types/services"
	"github.com/razorpay/statping/types/users"
	"github.com/razorpay/statping/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type SamplerFunc func() error

type Sampler interface {
	Samples() []database.DbObject
}

func TriggerSamples() error {
	return createSamples(
		core.Samples,
		//services.Samples,
		//messages.Samples,
		//checkins.Samples,
		//checkins.SamplesChkHits,
		//failures.Samples,
		//groups.Samples,
		//hits.Samples,
		//incidents.Samples,
	)
}

func createSamples(sm ...SamplerFunc) error {
	for _, v := range sm {
		if err := v(); err != nil {
			return err
		}
	}
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

// DropDatabase will DROP each table Statping created
func (d *DbConfig) DropDatabase() error {
	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}
	log.Infoln("Dropping Database Tables...")
	for _, t := range DbModels {
		if err := d.Db.DropTableIfExists(t); err != nil {
			return err.Error()
		}
		log.Infof("Dropped table: %T\n", t)
	}
	return nil
}

func (d *DbConfig) Close() {
	if d == nil {
		return
	}
	if d.Db != nil {
		d.Db.Close()
	}
}

// CreateDatabase will CREATE TABLES for each of the Statping elements
func (d *DbConfig) CreateDatabase() error {
	var err error

	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}

	log.Infoln("Creating Database Tables...")
	for _, table := range DbModels {
		if err := d.Db.CreateTable(table); err.Error() != nil {
			return errors.Wrap(err.Error(), fmt.Sprintf("error creating '%T' table", table))
		}
	}
	if err := d.Db.Table("core").CreateTable(&core.Core{}); err.Error() != nil {
		return errors.Wrap(err.Error(), fmt.Sprintf("error creating 'core' table"))
	}
	log.Infoln("Statping Database Created")

	return err
}
