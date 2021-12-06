package core

import (
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
)

func Example() *Core {
	core := &Core{
		Name:        "Statping Testing",
		Description: "This is a instance that runs for tests",
		ApiSecret:   "exampleapisecret",
		Domain:      "http://localhost:8080",
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		MigrationId: utils.Now().Unix(),
		Language:    "en",
	}
	App = core
	return App
}

func Samples() error {
	apiSecret := utils.Params.GetString("API_SECRET")

	if apiSecret == "" {
		apiSecret = utils.RandomString(32)
	}

	oauth := OAuth{Providers: "local"}

	core := &Core{
		Name:        utils.Params.GetString("NAME"),
		Description: utils.Params.GetString("DESCRIPTION"),
		ApiSecret:   apiSecret,
		Domain:      utils.Params.GetString("DOMAIN"),
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		MigrationId: utils.Now().Unix(),
		Language:    utils.Params.GetString("LANGUAGE"),
		OAuth:       oauth,
		Version:     utils.Params.GetString("VERSION"),
		Commit:      utils.Params.GetString("COMMIT"),
	}

	return core.Create()
}
