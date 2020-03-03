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
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/jinzhu/gorm"
)

var (
	// DbSession stores the Statping database session
	DbModels []interface{}
)

func init() {
	DbModels = []interface{}{&types.Service{}, &types.User{}, &types.Hit{}, &types.Failure{}, &types.Message{}, &types.Group{}, &types.Checkin{}, &types.CheckinHit{}, &notifier.Notification{}, &types.Incident{}, &types.IncidentUpdate{}, &types.Integration{}}
}

// CloseDB will close the database connection if available
func CloseDB() {
	database.Close()
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
func InsertCore(d *types.DbConfig) (*Core, error) {
	apiKey := utils.Getenv("API_KEY", utils.NewSHA1Hash(40))
	apiSecret := utils.Getenv("API_SECRET", utils.NewSHA1Hash(40))

	core := &types.Core{
		Name:        d.Project,
		Description: d.Description,
		ConfigFile:  "config.yml",
		ApiKey:      apiKey.(string),
		ApiSecret:   apiSecret.(string),
		Domain:      d.Domain,
		MigrationId: time.Now().Unix(),
	}

	CoreApp = &Core{Core: core}

	CoreApp.config = d

	_, err := database.Create(CoreApp.Core)
	return CoreApp, err
}

func findDbFile() string {
	if CoreApp.config == nil {
		return findSQLin(utils.Directory)
	}
	if CoreApp.config.SqlFile != "" {
		return CoreApp.config.SqlFile
	}
	return utils.Directory + "/" + types.SqliteFilename
}

func findSQLin(path string) string {
	filename := types.SqliteFilename
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".db" {
			fmt.Println("DB file is now: ", info.Name())
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
func (c *Core) Connect(configs *types.DbConfig, retry bool, location string) error {
	postgresSSL := os.Getenv("POSTGRES_SSLMODE")
	if database.Available() {
		return nil
	}
	var conn string
	var err error

	switch configs.DbConn {
	case "sqlite", "sqlite3":
		conn = findDbFile()
		configs.SqlFile = fmt.Sprintf("%s/%s", utils.Directory, conn)
		log.Infof("SQL database file at: %s", configs.SqlFile)
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
		log.Debugln(fmt.Sprintf("Database connection error %s", err))
		if retry {
			log.Errorln(fmt.Sprintf("Database %s connection to '%s' is not available, trying again in 5 seconds...", configs.DbConn, configs.DbHost))
			return c.waitForDb(configs)
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
		if utils.VerboseMode >= 4 {
			database.LogMode(true).Debug().SetLogger(gorm.Logger{log})
		}
		log.Infoln(fmt.Sprintf("Database %v connection was successful.", configs.DbConn))
	}

	CoreApp.config = configs

	return err
}

// waitForDb will sleep for 5 seconds and try to connect to the database again
func (c *Core) waitForDb(configs *types.DbConfig) error {
	time.Sleep(5 * time.Second)
	return c.Connect(configs, true, utils.Directory)
}

// Update will save the config.yml file
func (c *Core) UpdateConfig() error {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer config.Close()

	data, err := yaml.Marshal(c.config)
	if err != nil {
		log.Errorln(err)
		return err
	}
	config.WriteString(string(data))

	return err
}

// Save will initially create the config.yml file
func SaveConfig(d *types.DbConfig) error {
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer config.Close()
	log.WithFields(utils.ToFields(d)).Debugln("saving config file at: " + utils.Directory + "/config.yml")

	if d.ApiKey == "" || d.ApiSecret == "" {
		apiKey := utils.Getenv("API_KEY", utils.NewSHA1Hash(16))
		apiSecret := utils.Getenv("API_SECRET", utils.NewSHA1Hash(16))
		d.ApiKey = apiKey.(string)
		d.ApiSecret = apiSecret.(string)
	}
	if d.DbConn == "sqlite3" {
		d.DbConn = "sqlite"
	}

	data, err := yaml.Marshal(d)
	if err != nil {
		log.Errorln(err)
		return err
	}
	if _, err := config.WriteString(string(data)); err != nil {
		return errors.Wrap(err, "error writing to config.yml")
	}
	log.WithFields(utils.ToFields(d)).Infoln("Saved config file at: " + utils.Directory + "/config.yml")

	CoreApp.config = d
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
	_, err := database.Create(&newCore)
	if err == nil {
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
	for _, t := range DbModels {
		if err := database.Get().DropTableIfExists(t); err != nil {
			return err.Error()
		}
		log.Infof("Dropped table: %T\n", t)
	}
	return nil
}

// CreateDatabase will CREATE TABLES for each of the Statping elements
func (c *Core) CreateDatabase() error {
	var err error
	log.Infoln("Creating Database Tables...")
	for _, table := range DbModels {
		if err := database.Get().CreateTable(table); err.Error() != nil {
			return err.Error()
		}
	}
	if err := database.Get().Table("core").CreateTable(&types.Core{}); err.Error() != nil {
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

func (c *Core) ServicesFromEnvFile() error {
	servicesEnv := utils.Getenv("SERVICES_FILE", "").(string)
	if servicesEnv == "" {
		return nil
	}

	file, err := os.Open(servicesEnv)
	if err != nil {
		return errors.Wrapf(err, "error opening 'SERVICES_FILE' at: %s", servicesEnv)
	}
	defer file.Close()

	var serviceLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		serviceLines = append(serviceLines, scanner.Text())
	}

	if len(serviceLines) == 0 {
		return nil
	}

	for k, service := range serviceLines {

		svr, err := utils.ValidateService(service)
		if err != nil {
			return errors.Wrapf(err, "invalid service at index %d in SERVICES_FILE environment variable", k)
		}
		if findServiceByHash(svr.String()) == nil {
			if _, err := database.Create(svr); err != nil {
				return errors.Wrapf(err, "could not create service %s", svr.Name)
			}
			log.Infof("Created new service '%s'", svr.Name)
		}

	}

	return nil
}

// MigrateDatabase will migrate the database structure to current version.
// This function will NOT remove previous records, tables or columns from the database.
// If this function has an issue, it will ROLLBACK to the previous state.
func (c *Core) MigrateDatabase() error {
	log.Infoln("Migrating Database Tables...")
	tx := database.Begin("migration")
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

	if err := tx.Commit().Error(); err != nil {
		return err
	}
	log.Infoln("Statping Database Migrated")

	return nil
}
