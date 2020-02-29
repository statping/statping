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

package core

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

var (
	// DbSession stores the Statping database session
	DbSession database.Database
	DbModels  []interface{}
)

func init() {
	DbModels = []interface{}{&types.Service{}, &types.User{}, &types.Hit{}, &types.Failure{}, &types.Message{}, &types.Group{}, &types.Checkin{}, &types.CheckinHit{}, &notifier.Notification{}, &types.Incident{}, &types.IncidentUpdate{}, &types.Integration{}}
}

// DbConfig stores the config.yml file for the statup configuration
type DbConfig struct {
	*types.DbConfig
}

func Database(obj interface{}) database.Database {
	switch obj.(type) {
	case *types.Service, *Service, []*Service:
		return DbSession.Model(&types.Service{})
	case *types.Hit, *Hit, []*Hit:
		return DbSession.Model(&types.Hit{})
	case *types.Failure, *Failure, []*Failure:
		return DbSession.Model(&types.Failure{})
	case *types.Core, *Core:
		return DbSession.Table("core").Model(&CoreApp)
	case *types.Checkin, *Checkin, []*Checkin:
		return DbSession.Model(&types.Checkin{})
	case *types.CheckinHit, *CheckinHit, []*CheckinHit:
		return DbSession.Model(&types.CheckinHit{})
	case *types.User, *User, []*User:
		return DbSession.Model(&types.User{})
	case *types.Group, *Group, []*Group:
		return DbSession.Model(&types.Group{})
	case *types.Incident, *Incident, []*Incident:
		return DbSession.Model(&types.Incident{})
	case *types.IncidentUpdate, *IncidentUpdate, []*IncidentUpdate:
		return DbSession.Model(&types.IncidentUpdate{})
	case *types.Message, *Message, []*Message:
		return DbSession.Model(&types.Message{})
	default:
		return DbSession
	}
}

// CloseDB will close the database connection if available
func CloseDB() {
	if DbSession != nil {
		DbSession.DB().Close()
	}
}

//// AfterFind for Core will set the timezone
//func (c *Core) AfterFind() (err error) {
//	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
//	c.UpdatedAt = utils.Timezoner(c.UpdatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for Service will set the timezone
//func (s *Service) AfterFind() (err error) {
//	s.CreatedAt = utils.Timezoner(s.CreatedAt, CoreApp.Timezone)
//	s.UpdatedAt = utils.Timezoner(s.UpdatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for Hit will set the timezone
//func (h *Hit) AfterFind() (err error) {
//	h.CreatedAt = utils.Timezoner(h.CreatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for Failure will set the timezone
//func (f *Failure) AfterFind() (err error) {
//	f.CreatedAt = utils.Timezoner(f.CreatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for USer will set the timezone
//func (u *User) AfterFind() (err error) {
//	u.CreatedAt = utils.Timezoner(u.CreatedAt, CoreApp.Timezone)
//	u.UpdatedAt = utils.Timezoner(u.UpdatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for Checkin will set the timezone
//func (c *Checkin) AfterFind() (err error) {
//	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
//	c.UpdatedAt = utils.Timezoner(c.UpdatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for checkinHit will set the timezone
//func (c *CheckinHit) AfterFind() (err error) {
//	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
//	return
//}
//
//// AfterFind for Message will set the timezone
//func (u *Message) AfterFind() (err error) {
//	u.CreatedAt = utils.Timezoner(u.CreatedAt, CoreApp.Timezone)
//	u.UpdatedAt = utils.Timezoner(u.UpdatedAt, CoreApp.Timezone)
//	u.StartOn = utils.Timezoner(u.StartOn.UTC(), CoreApp.Timezone)
//	u.EndOn = utils.Timezoner(u.EndOn.UTC(), CoreApp.Timezone)
//	return
//}

// InsertCore create the single row for the Core settings in Statping
func (d *DbConfig) InsertCore() (*Core, error) {
	apiKey := utils.Getenv("API_KEY", utils.NewSHA1Hash(40))
	apiSecret := utils.Getenv("API_SECRET", utils.NewSHA1Hash(40))

	CoreApp = &Core{Core: &types.Core{
		Name:        d.Project,
		Description: d.Description,
		ConfigFile:  "config.yml",
		ApiKey:      apiKey.(string),
		ApiSecret:   apiSecret.(string),
		Domain:      d.Domain,
		MigrationId: time.Now().Unix(),
		Config:      d.DbConfig,
	}}
	query := DbSession.Create(CoreApp.Core)
	return CoreApp, query.Error()
}

func findDbFile() string {
	if CoreApp.Config.SqlFile != "" {
		return CoreApp.Config.SqlFile
	}
	filename := types.SqliteFilename
	err := filepath.Walk(utils.Directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".db" {
			filename = info.Name()
		}
		return nil
	})
	if err != nil {
		log.Error(err)
	}
	return filename
}

// Connect will attempt to connect to the sqlite, postgres, or mysql database
func (c *Core) Connect(retry bool, location string) error {
	postgresSSL := os.Getenv("POSTGRES_SSLMODE")
	if DbSession != nil {
		return nil
	}
	var conn string
	var err error

	if c.Config == nil {
		return errors.New("missing database connection configs")
	}

	configs := c.Config
	switch configs.DbConn {
	case "sqlite":
		sqlFilename := findDbFile()
		conn = sqlFilename
		log.Infof("SQL database file at: %v/%v", utils.Directory, conn)
		configs.DbConn = "sqlite3"
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
	log.WithFields(utils.ToFields(c, conn)).Debugln("attempting to connect to database")
	dbSession, err := database.Openw(configs.DbConn, conn)
	if err != nil {
		log.Debugln(fmt.Sprintf("Database connection error %v", err))
		if retry {
			log.Errorln(fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", configs.DbHost))
			return c.waitForDb()
		} else {
			return err
		}
	}
	log.WithFields(utils.ToFields(dbSession)).Debugln("connected to database")

	maxOpenConn := utils.Getenv("MAX_OPEN_CONN", 5)
	maxIdleConn := utils.Getenv("MAX_IDLE_CONN", 5)
	maxLifeConn := utils.Getenv("MAX_LIFE_CONN", 2*time.Minute)

	dbSession.DB().SetMaxOpenConns(maxOpenConn.(int))
	dbSession.DB().SetMaxIdleConns(maxIdleConn.(int))
	dbSession.DB().SetConnMaxLifetime(maxLifeConn.(time.Duration))

	if dbSession.DB().Ping() == nil {
		DbSession = dbSession
		if utils.VerboseMode >= 4 {
			DbSession.LogMode(true).Debug().SetLogger(gorm.Logger{log})
		}
		log.Infoln(fmt.Sprintf("Database %v connection was successful.", configs.DbConn))
	}
	return err
}

// waitForDb will sleep for 5 seconds and try to connect to the database again
func (c *Core) waitForDb() error {
	time.Sleep(5 * time.Second)
	return c.Connect(true, utils.Directory)
}

// Update will save the config.yml file
func (c *Core) UpdateConfig() error {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	data, err := yaml.Marshal(c.Config)
	if err != nil {
		log.Errorln(err)
		return err
	}
	config.WriteString(string(data))
	config.Close()
	return err
}

// Save will initially create the config.yml file
func (d *DbConfig) Save() error {
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer config.Close()
	log.WithFields(utils.ToFields(d)).Debugln("saving config file at: " + utils.Directory + "/config.yml")
	CoreApp.Config = d.DbConfig

	apiKey := utils.Getenv("API_KEY", utils.NewSHA1Hash(16))
	apiSecret := utils.Getenv("API_SECRET", utils.NewSHA1Hash(16))

	CoreApp.Config.ApiKey = apiKey.(string)
	CoreApp.Config.ApiSecret = apiSecret.(string)
	data, err := yaml.Marshal(d)
	if err != nil {
		log.Errorln(err)
		return err
	}
	config.WriteString(string(data))
	log.WithFields(utils.ToFields(d)).Infoln("saved config file at: " + utils.Directory + "/config.yml")
	return err
}

// CreateCore will initialize the global variable 'CoreApp". This global variable contains most of Statping app.
func (c *Core) CreateCore() *Core {
	newCore := &types.Core{
		Name:        c.Name,
		Description: c.Description,
		ConfigFile:  utils.Directory + "/config.yml",
		ApiKey:      c.ApiKey,
		ApiSecret:   c.ApiSecret,
		Domain:      c.Domain,
		MigrationId: time.Now().Unix(),
	}
	db := Database(newCore).Create(&newCore)
	if db.Error() == nil {
		CoreApp = &Core{Core: newCore}
	}
	CoreApp, err := SelectCore()
	if err != nil {
		log.Errorln(err)
	}
	return CoreApp
}

// DropDatabase will DROP each table Statping created
func (c *Core) DropDatabase() error {
	log.Infoln("Dropping Database Tables...")
	tables := []string{"checkins", "checkin_hits", "notifications", "core", "failures", "hits", "services", "users", "messages", "incidents", "incident_updates"}
	for _, t := range tables {
		if err := DbSession.DropTableIfExists(t); err != nil {
			return err.Error()
		}
	}
	return nil
}

// CreateDatabase will CREATE TABLES for each of the Statping elements
func (c *Core) CreateDatabase() error {
	var err error
	log.Infoln("Creating Database Tables...")
	for _, table := range DbModels {
		if err := DbSession.CreateTable(table); err.Error() != nil {
			return err.Error()
		}
	}
	if err := DbSession.Table("core").CreateTable(&types.Core{}); err.Error() != nil {
		return err.Error()
	}
	log.Infoln("Statping Database Created")
	return err
}

// findServiceByHas will return a service that matches the SHA256 hash of a service
// Service hash example: sha256(name:EXAMPLEdomain:HTTP://DOMAIN.COMport:8080type:HTTPmethod:GET)
func findServiceByHash(hash string) *Service {
	for _, service := range Services() {
		if service.String() == hash {
			return service
		}
	}
	return nil
}

func (c *Core) CreateServicesFromEnvs() error {
	servicesEnv := utils.Getenv("SERVICES", []*types.Service{}).([]*types.Service)

	for k, service := range servicesEnv {

		if err := service.Valid(); err != nil {
			return errors.Wrapf(err, "invalid service at index %d in SERVICES environment variable", k)
		}
		if findServiceByHash(service.String()) == nil {
			newService := &types.Service{
				Name:   service.Name,
				Domain: service.Domain,
				Method: service.Method,
				Type:   service.Type,
			}
			if _, err := database.Create(newService); err != nil {
				return errors.Wrapf(err, "could not create service %s", newService.Name)
			}
			log.Infof("Created new service '%s'", newService.Name)
		}

	}

	return nil
}

// MigrateDatabase will migrate the database structure to current version.
// This function will NOT remove previous records, tables or columns from the database.
// If this function has an issue, it will ROLLBACK to the previous state.
func (c *Core) MigrateDatabase() error {
	log.Infoln("Migrating Database Tables...")
	tx := DbSession.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error() != nil {
		log.Errorln(tx.Error())
		return tx.Error()
	}
	for _, table := range DbModels {
		tx = tx.AutoMigrate(table)
	}
	if err := tx.Table("core").AutoMigrate(&types.Core{}); err.Error() != nil {
		tx.Rollback()
		log.Errorln(fmt.Sprintf("Statping Database could not be migrated: %v", tx.Error()))
		return tx.Error()
	}
	log.Infoln("Statping Database Migrated")

	if err := tx.Commit().Error(); err != nil {
		return err
	}

	return nil
}
