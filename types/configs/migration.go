package configs

import (
	"fmt"
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
)

// InsertNotifierDB inject the Statping database instance to the Notifier package
//func (c *DbConfig) InsertNotifierDB() error {
//	if !database.Available() {
//		err := c.Connect()
//		if err != nil {
//			return errors.New("database connection has not been created")
//		}
//	}
//	notifiers.SetDB(database.DB())
//	return nil
//}

// InsertIntegratorDB inject the Statping database instance to the Integrations package
//func (c *DbConfig) InsertIntegratorDB() error {
//	if !database.Available() {
//		err := c.Connect()
//		if err != nil {
//			return errors.Wrap(err,"database connection has not been created")
//		}
//	}
//	integrations.SetDB(database.DB())
//	return nil
//}

//MigrateDatabase will migrate the database structure to current version.
//This function will NOT remove previous records, tables or columns from the database.
//If this function has an issue, it will ROLLBACK to the previous state.
func (c *DbConfig) MigrateDatabase() error {

	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}, &integrations.Integration{}}

	log.Infoln("Migrating Database Tables...")
	tx := database.Begin("migration")
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, table := range DbModels {
		tx = tx.AutoMigrate(table)
		if tx.Error() != nil {
			log.Errorln(tx.Error())
			return tx.Error()
		}
	}
	if err := tx.Table("core").AutoMigrate(&core.Core{}); err.Error() != nil {
		tx.Rollback()
		log.Errorln(fmt.Sprintf("Statping Database could not be migrated: %v", tx.Error()))
		return tx.Error()
	}

	if err := tx.Commit().Error(); err != nil {
		return err
	}
	log.Infoln("Statping Database Migrated")

	return nil
}
