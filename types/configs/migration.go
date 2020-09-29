package configs

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/utils"

	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
)

func (d *DbConfig) ResetCore() error {
	if d.Db.HasTable("core") {
		return nil
	}
	var srvs int64
	if d.Db.HasTable(&services.Service{}) {
		d.Db.Model(&services.Service{}).Count(&srvs)
		if srvs > 0 {
			return errors.New("there are already services setup.")
		}
	}
	if err := d.DropDatabase(); err != nil {
		return errors.Wrap(err, "error dropping database")
	}
	if err := d.CreateDatabase(); err != nil {
		return errors.Wrap(err, "error creating database")
	}
	if err := CreateAdminUser(); err != nil {
		return errors.Wrap(err, "error creating default admin user")
	}
	if utils.Params.GetBool("SAMPLE_DATA") {
		log.Infoln("Adding Sample Data")
		if err := TriggerSamples(); err != nil {
			return errors.Wrap(err, "error adding sample data")
		}
	} else {
		if err := core.Samples(); err != nil {
			return errors.Wrap(err, "error added core details")
		}
	}
	return nil
}

func (d *DbConfig) DatabaseChanges() error {
	var cr core.Core
	d.Db.Model(&core.Core{}).Find(&cr)

	if latestMigration > cr.MigrationId {
		log.Infof("Statping database is out of date, migrating to: %d", latestMigration)

		switch d.Db.DbType() {
		case "mysql":
			if err := d.genericMigration("MODIFY", false); err != nil {
				return err
			}
		case "postgres":
			if err := d.genericMigration("ALTER", true); err != nil {
				return err
			}
		default:
			if err := d.sqliteMigration(); err != nil {
				return err
			}
		}

		if err := d.Db.Exec(fmt.Sprintf("UPDATE core SET migration_id = %d", latestMigration)).Error(); err != nil {
			return err
		}

		if err := d.BackupAssets(); err != nil {
			return err
		}
	}
	return nil
}

// BackupAssets is a temporary function (to version 0.90.*) to backup your customized theme
// to a new folder called 'assets_backup'.
func (d *DbConfig) BackupAssets() error {
	if source.UsingAssets(utils.Directory) {
		log.Infof("Backing up 'assets' folder to 'assets_backup'")
		if err := utils.RenameDirectory(utils.Directory+"/assets", utils.Directory+"/assets_backup"); err != nil {
			return err
		}
		log.Infof("Old assets are now stored in: " + utils.Directory + "/assets_backup")
	}
	return nil
}

//MigrateDatabase will migrate the database structure to current version.
//This function will NOT remove previous records, tables or columns from the database.
//If this function has an issue, it will ROLLBACK to the previous state.
func (d *DbConfig) MigrateDatabase() error {
	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}

	log.Infoln("Migrating Database Tables...")
	tx := d.Db.Begin()
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

	log.Infof("Migrating App to version: %s (%s)", utils.Params.GetString("VERSION"), utils.Params.GetString("COMMIT"))
	if err := tx.Table("core").AutoMigrate(&core.Core{}); err.Error() != nil {
		tx.Rollback()
		log.Errorln(fmt.Sprintf("Statping Database could not be migrated: %v", tx.Error()))
		return tx.Error()
	}

	if err := tx.Commit().Error(); err != nil {
		return err
	}

	d.Db.Table("core").Model(&core.Core{}).Update("version", utils.Params.GetString("VERSION"))

	log.Infoln("Statping Database Tables Migrated")

	if err := d.Db.Model(&hits.Hit{}).AddIndex("idx_service_hit", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&hits.Hit{}).AddIndex("hit_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("fail_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("idx_service_fail", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("idx_checkin_fail", "checkin").Error(); err != nil {
		log.Errorln(err)
	}
	log.Infoln("Database Indexes Created")

	return nil
}
