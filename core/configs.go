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
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"io/ioutil"
	"os"
)

// ErrorResponse is used for HTTP errors to show to User
type ErrorResponse struct {
	Error string
}

// LoadConfigFile will attempt to load the 'config.yml' file in a specific directory
func LoadConfigFile(directory string) (*types.DbConfig, error) {
	var configs *types.DbConfig
	if os.Getenv("DB_CONN") != "" {
		log.Warnln("DB_CONN environment variable was found, waiting for database...")
		return LoadUsingEnv()
	}
	log.Debugln("attempting to read config file at: " + directory + "/config.yml")
	file, err := ioutil.ReadFile(directory + "/config.yml")
	if err != nil {
		return nil, errors.New("config.yml file not found at " + directory + "/config.yml - starting in setup mode")
	}
	err = yaml.Unmarshal(file, &configs)
	if err != nil {
		return nil, err
	}
	log.WithFields(utils.ToFields(configs)).Debugln("read config file: " + directory + "/config.yml")
	CoreApp.Config = configs
	return configs, err
}

// LoadUsingEnv will attempt to load database configs based on environment variables. If DB_CONN is set if will force this function.
func LoadUsingEnv() (*types.DbConfig, error) {
	Configs, err := EnvToConfig()
	if err != nil {
		return Configs, err
	}
	CoreApp.Name = os.Getenv("NAME")
	if Configs.Domain == "" {
		CoreApp.Domain = Configs.LocalIP
	} else {
		CoreApp.Domain = os.Getenv("DOMAIN")
	}
	CoreApp.UseCdn = types.NewNullBool(os.Getenv("USE_CDN") == "true")

	err = CoreApp.Connect(true, utils.Directory)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	CoreApp.SaveConfig(Configs)
	exists := DbSession.HasTable("core")
	if !exists {
		log.Infoln(fmt.Sprintf("Core database does not exist, creating now!"))
		CoreApp.DropDatabase()
		CoreApp.CreateDatabase()
		CoreApp, err = CoreApp.InsertCore(Configs)
		if err != nil {
			log.Errorln(err)
		}

		username := os.Getenv("ADMIN_USER")
		if username == "" {
			username = "admin"
		}
		password := os.Getenv("ADMIN_PASSWORD")
		if password == "" {
			password = "admin"
		}

		admin := ReturnUser(&types.User{
			Username: username,
			Password: password,
			Email:    "info@admin.com",
			Admin:    types.NewNullBool(true),
		})
		_, err := admin.Create()

		SampleData()
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
func EnvToConfig() (*types.DbConfig, error) {
	var err error
	if os.Getenv("DB_CONN") == "" {
		return nil, errors.New("Missing DB_CONN environment variable")
	}
	if os.Getenv("DB_CONN") != "sqlite" {
		if os.Getenv("DB_HOST") == "" {
			return nil, errors.New("Missing DB_HOST environment variable")
		}
		if os.Getenv("DB_USER") == "" {
			return nil, errors.New("Missing DB_USER environment variable")
		}
		if os.Getenv("DB_PASS") == "" {
			return nil, errors.New("Missing DB_PASS environment variable")
		}
		if os.Getenv("DB_DATABASE") == "" {
			return nil, errors.New("Missing DB_DATABASE environment variable")
		}
	}
	port := utils.ToInt(os.Getenv("DB_PORT"))
	if port == 0 {
		port = defaultPort(os.Getenv("DB_PORT"))
	}
	name := os.Getenv("NAME")
	if name == "" {
		name = "Statping"
	}
	description := os.Getenv("DESCRIPTION")
	if description == "" {
		description = "Statping Monitoring Sample Data"
	}

	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		adminUser = "admin"
	}

	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		adminPass = "admin"
	}

	configs := &types.DbConfig{
		DbConn:      os.Getenv("DB_CONN"),
		DbHost:      os.Getenv("DB_HOST"),
		DbUser:      os.Getenv("DB_USER"),
		DbPass:      os.Getenv("DB_PASS"),
		DbData:      os.Getenv("DB_DATABASE"),
		DbPort:      port,
		Project:     name,
		Description: description,
		Domain:      os.Getenv("DOMAIN"),
		Email:       "",
		Username:    adminUser,
		Password:    adminPass,
		Error:       nil,
		Location:    utils.Directory,
		SqlFile:     os.Getenv("SQL_FILE"),
	}
	CoreApp.Config = configs
	return configs, err
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
