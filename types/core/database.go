package core

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Core{})
	c, err := Select()
	if err != nil {
		utils.Log.Errorln(err)
		return
	}
	apiEnv := utils.Params.GetString("API_SECRET")
	if c.ApiSecret != apiEnv && apiEnv != "" {
		c.ApiSecret = apiEnv
		if err := c.Update(); err != nil {
			utils.Log.Errorln(err)
		}
	}
}

func (c *Core) AfterFind() {
	metrics.Query("core", "find")
}

func Select() (*Core, error) {
	var c Core
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
	if utils.Params.GetString("LANGUAGE") != "" {
		App.Language = utils.Params.GetString("LANGUAGE")
	}
	if utils.Params.GetString("API_SECRET") != "" {
		App.ApiSecret = utils.Params.GetString("API_SECRET")
	}
	App.Version = utils.Params.GetString("VERSION")
	App.Commit = utils.Params.GetString("COMMIT")
	return App, q.Error()
}

func (c *Core) Create() error {
	if c.ApiSecret == "" {
		c.ApiSecret = utils.RandomString(16)
		apiEnv := utils.Params.GetString("API_SECRET")
		if apiEnv != "" {
			c.ApiSecret = apiEnv
		}
	}
	q := db.Create(c)
	utils.Log.Infof("API Key created: %s", c.ApiSecret)
	return q.Error()
}

func (c *Core) Update() error {
	q := db.UpdateColumns(c)
	return q.Error()
}

func (c *Core) Delete() error {
	return nil
}
