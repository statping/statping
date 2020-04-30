package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"time"
)

// Connect will attempt to connect to the sqlite, postgres, or mysql database
func Connect(configs *DbConfig, retry bool) error {
	conn := configs.ConnectionString()
	p := utils.Params
	var err error

	log.WithFields(utils.ToFields(configs, conn)).Debugln("attempting to connect to database")

	dbSession, err := database.Openw(configs.DbConn, conn)
	if err != nil {
		log.Errorf(fmt.Sprintf("Database connection error %s", err))
		if retry {
			log.Warnln(fmt.Sprintf("Database %s connection to '%s' is not available, trying again in 5 seconds...", configs.DbConn, configs.DbHost))
			time.Sleep(5 * time.Second)
			return Connect(configs, retry)
		} else {
			return err
		}
	}

	apiSecret := p.GetString("API_SECRET")
	configs.ApiSecret = apiSecret

	log.WithFields(utils.ToFields(dbSession)).Debugln("connected to database")

	maxOpenConn := p.GetInt("MAX_OPEN_CONN")
	maxIdleConn := p.GetInt("MAX_IDLE_CONN")
	maxLifeConn := p.GetDuration("MAX_LIFE_CONN")

	dbSession.DB().SetMaxOpenConns(maxOpenConn)
	dbSession.DB().SetMaxIdleConns(maxIdleConn)
	dbSession.DB().SetConnMaxLifetime(maxLifeConn)

	if dbSession.DB().Ping() == nil {
		if utils.VerboseMode >= 4 {
			dbSession.LogMode(true).Debug().SetLogger(gorm.Logger{log})
		}
		log.Infoln(fmt.Sprintf("Database %s connection was successful.", configs.DbConn))
	}

	configs.Db = dbSession

	initModels(configs.Db)

	return err
}

func initModels(db database.Database) {
	core.SetDB(db)
	services.SetDB(db)
	hits.SetDB(db)
	failures.SetDB(db)
	checkins.SetDB(db)
	notifications.SetDB(db)
	incidents.SetDB(db)
	users.SetDB(db)
	messages.SetDB(db)
	groups.SetDB(db)
}

func CreateAdminUser(c *DbConfig) error {
	log.Infoln(fmt.Sprintf("Default Admininstrator user does not exist, creating now! (admin/admin)"))

	adminUser := utils.Params.GetString("ADMIN_USER")
	adminPass := utils.Params.GetString("ADMIN_PASSWORD")

	if adminUser == "" || adminPass == "" {
		adminUser = "admin"
		adminPass = "admin"
	}

	admin := &users.User{
		Username: adminUser,
		Password: adminPass,
		Email:    "info@admin.com",
		Admin:    null.NewNullBool(true),
	}

	if err := admin.Create(); err != nil {
		return errors.Wrap(err, "error creating admin")
	}

	return nil
}
