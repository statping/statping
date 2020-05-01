package core

import (
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
)

func Samples() error {
	apiSecret := utils.Params.GetString("API_SECRET")

	if apiSecret == "" {
		apiSecret = utils.RandomString(32)
	}

	core := &Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		ApiSecret:   apiSecret,
		Domain:      "http://localhost:8080",
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		MigrationId: utils.Now().Unix(),
	}

	return core.Create()
}
