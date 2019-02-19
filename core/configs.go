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
func LoadConfigFile(directory string) (*DbConfig, error) {
	var configs *DbConfig
	if os.Getenv("DB_CONN") != "" {
		utils.Log(1, "DB_CONN environment variable was found, waiting for database...")
		return LoadUsingEnv()
	}
	file, err := ioutil.ReadFile(directory + "/config.yml")
	if err != nil {
		return nil, errors.New("config.yml file not found at " + directory + "/config.yml - starting in setup mode")
	}
	err = yaml.Unmarshal(file, &configs)
	if err != nil {
		return nil, err
	}
	Configs = configs
	return Configs, err
}

// LoadUsingEnv will attempt to load database configs based on environment variables. If DB_CONN is set if will force this function.
func LoadUsingEnv() (*DbConfig, error) {
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
	CoreApp.DbConnection = Configs.DbConn
	CoreApp.UseCdn = types.NewNullBool(os.Getenv("USE_CDN") == "true")

	err = Configs.Connect(true, utils.Directory)
	if err != nil {
		utils.Log(4, err)
		return nil, err
	}
	Configs.Save()
	exists := DbSession.HasTable("core")
	if !exists {
		utils.Log(1, fmt.Sprintf("Core database does not exist, creating now!"))
		Configs.DropDatabase()
		Configs.CreateDatabase()
		CoreApp, err = Configs.InsertCore()
		if err != nil {
			utils.Log(3, err)
		}

		admin := ReturnUser(&types.User{
			Username: "admin",
			Password: "admin",
			Email:    "info@admin.com",
			Admin:    types.NewNullBool(true),
		})
		_, err := admin.Create()

		SampleData()
		return Configs, err
	}
	return Configs, nil
}

// DefaultPort accepts a database type and returns its default port
func DefaultPort(db string) int64 {
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
	Configs = new(DbConfig)
	var err error
	if os.Getenv("DB_CONN") == "" {
		return Configs, errors.New("Missing DB_CONN environment variable")
	}
	if os.Getenv("DB_CONN") != "sqlite" {
		if os.Getenv("DB_HOST") == "" {
			return Configs, errors.New("Missing DB_HOST environment variable")
		}
		if os.Getenv("DB_USER") == "" {
			return Configs, errors.New("Missing DB_USER environment variable")
		}
		if os.Getenv("DB_PASS") == "" {
			return Configs, errors.New("Missing DB_PASS environment variable")
		}
		if os.Getenv("DB_DATABASE") == "" {
			return Configs, errors.New("Missing DB_DATABASE environment variable")
		}
	}
	port := utils.ToInt(os.Getenv("DB_PORT"))
	if port == 0 {
		port = DefaultPort(os.Getenv("DB_PORT"))
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

	Configs = &DbConfig{
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
	}
	return Configs, err
}

// SampleData runs all the sample data for a new Statping installation
func SampleData() error {
	if err := InsertSampleData(); err != nil {
		return err
	}
	if err := InsertSampleHits(); err != nil {
		return err
	}
	return nil
}

// DeleteConfig will delete the 'config.yml' file
func DeleteConfig() error {
	err := os.Remove(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(3, err)
		return err
	}
	return nil
}
