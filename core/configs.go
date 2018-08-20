// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"os"
	"time"
)

func LoadConfig() (*types.Config, error) {
	if os.Getenv("DB_CONN") != "" {
		utils.Log(1, "DB_CONN environment variable was found, waiting for database...")
		return LoadUsingEnv()
	}
	Configs = new(types.Config)
	file, err := ioutil.ReadFile(utils.Directory + "/config.yml")
	if err != nil {
		return nil, errors.New("config.yml file not found - starting in setup mode")
	}
	err = yaml.Unmarshal(file, &Configs)
	if err != nil {
		return nil, err
	}
	CoreApp.DbConnection = Configs.Connection
	return Configs, err
}

func LoadUsingEnv() (*types.Config, error) {
	Configs = new(types.Config)
	if os.Getenv("DB_CONN") == "" {
		return nil, errors.New("Missing DB_CONN environment variable")
	}
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
	Configs.Connection = os.Getenv("DB_CONN")
	Configs.Host = os.Getenv("DB_HOST")
	Configs.Port = os.Getenv("DB_PORT")
	Configs.User = os.Getenv("DB_USER")
	Configs.Password = os.Getenv("DB_PASS")
	Configs.Database = os.Getenv("DB_DATABASE")
	CoreApp.DbConnection = os.Getenv("DB_CONN")
	CoreApp.Name = os.Getenv("NAME")
	CoreApp.Domain = os.Getenv("DOMAIN")
	if os.Getenv("USE_CDN") == "true" {
		CoreApp.UseCdn = true
	}

	dbConfig := &DbConfig{&types.DbConfig{
		DbConn:      os.Getenv("DB_CONN"),
		DbHost:      os.Getenv("DB_HOST"),
		DbUser:      os.Getenv("DB_USER"),
		DbPass:      os.Getenv("DB_PASS"),
		DbData:      os.Getenv("DB_DATABASE"),
		DbPort:      5432,
		Project:     "Statup - " + os.Getenv("NAME"),
		Description: "New Statup Installation",
		Domain:      os.Getenv("DOMAIN"),
		Username:    "admin",
		Password:    "admin",
		Email:       "info@localhost.com",
	}}

	err := DbConnection(dbConfig.DbConn, true, utils.Directory)
	if err != nil {
		utils.Log(4, err)
		return nil, err
	}

	exists, err := DbSession.Collection("core").Find().Exists()
	if !exists {

		utils.Log(1, fmt.Sprintf("Core database does not exist, creating now!"))
		DropDatabase()
		CreateDatabase()

		CoreApp = &Core{Core: &types.Core{
			Name:        dbConfig.Project,
			Description: dbConfig.Description,
			Config:      "config.yml",
			ApiKey:      utils.NewSHA1Hash(9),
			ApiSecret:   utils.NewSHA1Hash(16),
			Domain:      dbConfig.Domain,
			MigrationId: time.Now().Unix(),
		}}

		CoreApp.DbConnection = dbConfig.DbConn

		err := InsertCore(CoreApp)
		if err != nil {
			utils.Log(3, err)
		}

		admin := &types.User{
			Username: "admin",
			Password: "admin",
			Email:    "info@admin.com",
			Admin:    true,
		}
		CreateUser(admin)

		LoadSampleData()

		return Configs, err

	}

	return Configs, nil
}
