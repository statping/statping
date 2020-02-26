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
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
)

// ErrorResponse is used for HTTP errors to show to User
type ErrorResponse struct {
	Error string
}

// LoadConfigFile will attempt to load the 'config.yml' file in a specific directory
func LoadConfigFile(directory string) (*DbConfig, error) {
	var configs *DbConfig

	dbConn := utils.Getenv("DB_CONN", "")

	if dbConn != "" {
		log.Infof("DB_CONN=%s environment variable was found, waiting for database...", dbConn)
		return LoadUsingEnv()
	}
	log.Debugln("Attempting to read config file at: " + directory + "/config.yml")
	file, err := utils.OpenFile(directory + "/config.yml")
	if err != nil {
		CoreApp.Setup = false
		return nil, errors.New("config.yml file not found at " + directory + "/config.yml - starting in setup mode")
	}
	err = yaml.Unmarshal([]byte(file), &configs)
	if err != nil {
		return nil, err
	}
	log.WithFields(utils.ToFields(configs)).Debugln("read config file: " + directory + "/config.yml")
	CoreApp.Config = configs.DbConfig
	return configs, err
}

// LoadUsingEnv will attempt to load database configs based on environment variables. If DB_CONN is set if will force this function.
func LoadUsingEnv() (*DbConfig, error) {
	Configs, err := EnvToConfig()
	if err != nil {
		return Configs, err
	}

	CoreApp.Name = utils.Getenv("NAME", "").(string)
	CoreApp.Domain = utils.Getenv("DOMAIN", Configs.LocalIP).(string)
	CoreApp.UseCdn = types.NewNullBool(utils.Getenv("USE_CDN", false).(bool))

	err = CoreApp.Connect(true, utils.Directory)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	if err := Configs.Save(); err != nil {
		return nil, err
	}
	exists := DbSession.HasTable("core")
	if !exists {
		log.Infoln(fmt.Sprintf("Core database does not exist, creating now!"))
		if err := CoreApp.DropDatabase(); err != nil {
			return nil, err
		}
		if err := CoreApp.CreateDatabase(); err != nil {
			return nil, err
		}
		CoreApp, err = Configs.InsertCore()
		if err != nil {
			log.Errorln(err)
		}

		username := utils.Getenv("ADMIN_USER", "admin").(string)
		password := utils.Getenv("ADMIN_PASSWORD", "admin").(string)

		admin := &types.User{
			Username: username,
			Password: utils.HashPassword(password),
			Email:    "info@admin.com",
			Admin:    types.NewNullBool(true),
		}
		if _, err := database.Create(admin); err != nil {
			return nil, err
		}

		if err := SampleData(); err != nil {
			return nil, err
		}

		return Configs, err
	}
	return Configs, nil
}

// defaultPort accepts a database type and returns its default port
func defaultPort(db string) int64 {
	switch db {
	case "mysql":
		return 3306
	case "postgres":
		return 5432
	case "mssql":
		return 1433
	default:
		return 0
	}
}

// EnvToConfig converts environment variables to a DbConfig variable
func EnvToConfig() (*DbConfig, error) {
	var err error

	dbConn := utils.Getenv("DB_CONN", "").(string)
	dbHost := utils.Getenv("DB_HOST", "").(string)
	dbUser := utils.Getenv("DB_USER", "").(string)
	dbPass := utils.Getenv("DB_PASS", "").(string)
	dbData := utils.Getenv("DB_DATABASE", "").(string)
	dbPort := utils.Getenv("DB_PORT", defaultPort(dbConn)).(int64)
	name := utils.Getenv("NAME", "Statping").(string)
	desc := utils.Getenv("DESCRIPTION", "Statping Monitoring Sample Data").(string)
	user := utils.Getenv("ADMIN_USER", "admin").(string)
	password := utils.Getenv("ADMIN_PASS", "admin").(string)
	domain := utils.Getenv("DOMAIN", "").(string)
	sqlFile := utils.Getenv("SQL_FILE", "").(string)

	if dbConn != "sqlite" {
		if dbHost == "" {
			return nil, errors.New("Missing DB_HOST environment variable")
		}
		if dbUser == "" {
			return nil, errors.New("Missing DB_USER environment variable")
		}
		if dbPass == "" {
			return nil, errors.New("Missing DB_PASS environment variable")
		}
		if dbData == "" {
			return nil, errors.New("Missing DB_DATABASE environment variable")
		}
	}

	CoreApp.Config = &types.DbConfig{
		DbConn:      dbConn,
		DbHost:      dbHost,
		DbUser:      dbUser,
		DbPass:      dbPass,
		DbData:      dbData,
		DbPort:      dbPort,
		Project:     name,
		Description: desc,
		Domain:      domain,
		Email:       "",
		Username:    user,
		Password:    password,
		Error:       nil,
		Location:    utils.Directory,
		SqlFile:     sqlFile,
	}

	return &DbConfig{CoreApp.Config}, err
}

// SampleData runs all the sample data for a new Statping installation
func SampleData() error {
	if err := InsertSampleData(); err != nil {
		log.Errorln(err)
		return err
	}
	if err := InsertSampleHits(); err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

// DeleteConfig will delete the 'config.yml' file
func DeleteConfig() error {
	log.Debugln("deleting config yaml file", utils.Directory+"/config.yml")
	err := utils.DeleteFile(utils.Directory + "/config.yml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}
