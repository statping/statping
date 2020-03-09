package configs

import (
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/types/users"
	"github.com/hunterlong/statping/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"os"
	"time"
)

// Connect will attempt to connect to the sqlite, postgres, or mysql database
func Connect(configs *DbConfig, retry bool) error {
	postgresSSL := os.Getenv("POSTGRES_SSLMODE")
	if database.Available() {
		return nil
	}
	var conn string
	var err error

	switch configs.DbConn {
	case "sqlite", "sqlite3", "memory":
		if configs.DbConn == "memory" {
			conn = "sqlite3"
			configs.DbConn = ":memory"
		} else {
			conn = findDbFile(configs)
			configs.SqlFile = conn
			log.Infof("SQL database file at: %s", configs.SqlFile)
			configs.DbConn = "sqlite3"
		}
	case "mysql":
		host := fmt.Sprintf("%v:%v", configs.DbHost, configs.DbPort)
		conn = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC&time_zone=%%27UTC%%27", configs.DbUser, configs.DbPass, host, configs.DbData)
	case "postgres":
		sslMode := "disable"
		if postgresSSL != "" {
			sslMode = postgresSSL
		}
		conn = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v timezone=UTC sslmode=%v", configs.DbHost, configs.DbPort, configs.DbUser, configs.DbData, configs.DbPass, sslMode)
	case "mssql":
		host := fmt.Sprintf("%v:%v", configs.DbHost, configs.DbPort)
		conn = fmt.Sprintf("sqlserver://%v:%v@%v?database=%v", configs.DbUser, configs.DbPass, host, configs.DbData)
	}
	log.WithFields(utils.ToFields(configs, conn)).Debugln("attempting to connect to database")

	dbSession, err := database.Openw(configs.DbConn, conn)
	if err != nil {
		log.Debugln(fmt.Sprintf("Database connection error %s", err))
		if retry {
			log.Errorln(fmt.Sprintf("Database %s connection to '%s' is not available, trying again in 5 seconds...", configs.DbConn, configs.DbHost))
			time.Sleep(5 * time.Second)
			return Connect(configs, retry)
		} else {
			return err
		}
	}

	apiKey := utils.Getenv("API_KEY", utils.RandomString(16)).(string)
	apiSecret := utils.Getenv("API_SECRET", utils.RandomString(16)).(string)
	configs.ApiKey = apiKey
	configs.ApiSecret = apiSecret

	log.WithFields(utils.ToFields(dbSession)).Debugln("connected to database")

	maxOpenConn := utils.Getenv("MAX_OPEN_CONN", 5)
	maxIdleConn := utils.Getenv("MAX_IDLE_CONN", 5)
	maxLifeConn := utils.Getenv("MAX_LIFE_CONN", 2*time.Minute)

	dbSession.DB().SetMaxOpenConns(maxOpenConn.(int))
	dbSession.DB().SetMaxIdleConns(maxIdleConn.(int))
	dbSession.DB().SetConnMaxLifetime(maxLifeConn.(time.Duration))

	if dbSession.DB().Ping() == nil {
		if utils.VerboseMode >= 4 {
			database.LogMode(true).Debug().SetLogger(gorm.Logger{log})
		}
		log.Infoln(fmt.Sprintf("Database %v connection was successful.", configs.DbConn))
	}

	return err
}

func CreateAdminUser(configs *DbConfig) error {
	log.Infoln(fmt.Sprintf("Core database does not exist, creating now!"))

	if configs.Username == "" && configs.Password == "" {
		configs.Username = utils.Getenv("ADMIN_USER", "admin").(string)
		configs.Password = utils.Getenv("ADMIN_PASSWORD", "admin").(string)
	}

	admin := &users.User{
		Username: configs.Username,
		Password: configs.Password,
		Email:    "info@admin.com",
		Admin:    null.NewNullBool(true),
	}

	if err := admin.Create(); err != nil {
		return errors.Wrap(err, "error creating admin")
	}

	return nil
}
