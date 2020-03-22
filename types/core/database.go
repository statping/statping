package core

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"os"
	"time"
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
	App.UseCdn = null.NewNullBool(os.Getenv("USE_CDN") == "true")
	return App, q.Error()
}

func (c *Core) Create() error {
	c.ApiKey = c.ApiKey
	c.ApiSecret = c.ApiSecret
	apiKey := utils.Getenv("API_KEY", utils.NewSHA256Hash()).(string)
	apiSecret := utils.Getenv("API_SECRET", utils.NewSHA256Hash()).(string)

	if c.ApiKey == "" || c.ApiSecret == "" {
		c.ApiSecret = apiSecret
		c.ApiKey = apiKey
	}
	newCore := &Core{
		Name:        c.Name,
		Description: c.Description,
		ConfigFile:  utils.Directory + "/config.yml",
		ApiKey:      c.ApiKey,
		ApiSecret:   c.ApiSecret,
		Version:     App.Version,
		Domain:      c.Domain,
		MigrationId: time.Now().Unix(),
	}
	q := db.Create(&newCore)
	return q.Error()
}

func (c *Core) Update() error {
	q := db.Update(c)
	return q.Error()
}

func (c *Core) Delete() error {
	return nil
}

func Sample() error {
	apiKey := utils.Getenv("API_KEY", "samplekey")
	apiSecret := utils.Getenv("API_SECRET", "samplesecret")

	core := &Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		ApiKey:      apiKey.(string),
		ApiSecret:   apiSecret.(string),
		Domain:      "http://localhost:8080",
		Version:     "test",
		CreatedAt:   time.Now().UTC(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
	}

	q := db.Create(core)
	return q.Error()
}
