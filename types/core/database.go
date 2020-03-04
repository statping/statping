package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
	"os"
	"time"
)

func DB() database.Database {
	return database.DB().Table("core")
}

func Select() (*Core, error) {
	var c Core
	// SelectCore will return the CoreApp global variable and the settings/configs for Statping
	if !database.Available() {
		return nil, errors.New("database has not been initiated yet.")
	}
	exists := database.DB().HasTable("core")
	if !exists {
		return nil, errors.New("core database has not been setup yet.")
	}
	db := database.DB().Find(&c).Debug()
	if db.Error() != nil {
		return nil, db.Error()
	}
	App = &c
	App.UseCdn = null.NewNullBool(os.Getenv("USE_CDN") == "true")
	return App, db.Error()

}

func (c *Core) Create() error {
	//apiKey := utils.Getenv("API_KEY", utils.NewSHA1Hash(40))
	//apiSecret := utils.Getenv("API_SECRET", utils.NewSHA1Hash(40))
	newCore := &Core{
		Name:        c.Name,
		Description: c.Description,
		ConfigFile:  utils.Directory + "/config.yml",
		ApiKey:      c.ApiKey,
		ApiSecret:   c.ApiSecret,
		Domain:      c.Domain,
		MigrationId: time.Now().Unix(),
	}
	db := DB().FirstOrCreate(&newCore)
	return db.Error()
}

func (c *Core) Update() error {
	db := DB().Update(&c)
	return db.Error()
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

	db := database.DB().Create(&core)
	return db.Error()
}
