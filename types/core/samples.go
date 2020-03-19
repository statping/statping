package core

import (
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
)

func Samples() error {
	apiKey := utils.Getenv("API_KEY", utils.RandomString(16))
	apiSecret := utils.Getenv("API_SECRET", utils.RandomString(16))

	core := &Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		ApiKey:      apiKey.(string),
		ApiSecret:   apiSecret.(string),
		Domain:      "http://localhost:8080",
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		MigrationId: utils.Now().Unix(),
	}

	return core.Create()
}
