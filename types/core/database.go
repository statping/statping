package core

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Core{})
}

func Select() (*Core, error) {
	var c Core
	// SelectCore will return the CoreApp global variable and the settings/configs for Statping
	if err := db.DB().Ping(); err != nil {
		return nil, errors.New("database has not been initiated yet.")
	}
	exists := db.HasTable("core")
	if !exists {
		return nil, errors.New("core database has not been setup yet.")
	}
	q := db.Find(&c)
	if q.Error() != nil {
		return nil, db.Error()
	}
	App = &c

	if utils.Params.GetBool("USE_CDN") {
		App.UseCdn = null.NewNullBool(true)
	}
	if utils.Params.GetBool("ALLOW_REPORTS") {
		App.AllowReports = null.NewNullBool(true)
	}
	return App, q.Error()
}

func (c *Core) Create() error {
	secret := utils.Params.GetString("API_SECRET")
	if secret == "" {
		secret = utils.RandomString(32)
	}
	newCore := &Core{
		Name:        c.Name,
		Description: c.Description,
		ConfigFile:  utils.Directory + "/config.yml",
		ApiSecret:   secret,
		Version:     App.Version,
		Domain:      c.Domain,
		MigrationId: utils.Now().Unix(),
	}
	q := db.Create(&newCore)
	utils.Log.Infof("API Key created: %s", secret)
	return q.Error()
}

func (c *Core) Update() error {
	q := db.UpdateColumns(c)
	return q.Error()
}

func (c *Core) Delete() error {
	return nil
}
