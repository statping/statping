package configs

import (
	"fmt"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func (c *DbConfig) VerifyMigration() error {

	query := `
BEGIN TRANSACTION;
ALTER TABLE hits ALTER COLUMN latency BIGINT;
ALTER TABLE hits ALTER COLUMN ping_time BIGINT; 
ALTER TABLE failures ALTER COLUMN ping_time BIGINT;
UPDATE hits SET latency = CAST(latency * 10000 AS BIGINT);
UPDATE hits SET ping_time = CAST(ping_time * 100000 AS BIGINT);
UPDATE failures SET ping_time = CAST(ping_time * 100000 AS BIGINT);
COMMIT;`

	fmt.Println(c.Db.DbType())

	q := c.Db.Raw(query).Debug()

	return q.Error()
}

//MigrateDatabase will migrate the database structure to current version.
//This function will NOT remove previous records, tables or columns from the database.
//If this function has an issue, it will ROLLBACK to the previous state.
func (c *DbConfig) MigrateDatabase() error {

	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}

	log.Infoln("Migrating Database Tables...")
	tx := c.Db.Begin()
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
	log.Infoln("Statping Database Tables Migrated")

	if err := c.Db.Model(&hits.Hit{}).AddIndex("idx_service_hit", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := c.Db.Model(&hits.Hit{}).AddIndex("hit_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := c.Db.Model(&failures.Failure{}).AddIndex("fail_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := c.Db.Model(&failures.Failure{}).AddIndex("idx_service_fail", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := c.Db.Model(&failures.Failure{}).AddIndex("idx_checkin_fail", "checkin").Error(); err != nil {
		log.Errorln(err)
	}
	log.Infoln("Database Indexes Created")

	return nil
}
