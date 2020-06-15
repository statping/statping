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
		Name:        utils.Params.GetString("NAME"),
		Description: utils.Params.GetString("DESCRIPTION"),
		ApiSecret:   apiSecret,
		Domain:      utils.Params.GetString("DOMAIN"),
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		MigrationId: utils.Now().Unix(),
		Language:    utils.Params.GetString("LANGUAGE"),
	}

	return core.Create()
}
