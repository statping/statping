package configs

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func (c *DbConfig) DatabaseChanges() error {
	var cr core.Core
	c.Db.Model(&core.Core{}).Find(&cr)

	if latestMigration > cr.MigrationId {
		log.Infof("Statping database is out of date, migrating to: %d", latestMigration)

		switch c.Db.DbType() {
		case "mysql":
			if err := c.genericMigration("MODIFY", false); err != nil {
				return err
			}
		case "postgres":
			if err := c.genericMigration("ALTER", true); err != nil {
				return err
			}
		default:
			if err := c.sqliteMigration(); err != nil {
				return err
			}
		}

		if err := c.Db.Exec(fmt.Sprintf("UPDATE core SET migration_id = %d", latestMigration)).Error(); err != nil {
			return err
		}

		if err := c.BackupAssets(); err != nil {
			return err
		}
	}
	return nil
}

// BackupAssets is a temporary function (to version 0.90.*) to backup your customized theme
// to a new folder called 'assets_backup'.
func (c *DbConfig) BackupAssets() error {
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

	log.Infof("Migrating App to version: %s", core.App.Version)
	if err := tx.Table("core").AutoMigrate(&core.Core{}); err.Error() != nil {
		tx.Rollback()
		log.Errorln(fmt.Sprintf("Statping Database could not be migrated: %v", tx.Error()))
		return tx.Error()
	}

	if err := tx.Commit().Error(); err != nil {
		return err
	}

	c.Db.Table("core").Model(&core.Core{}).Update("version", core.App.Version)

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
