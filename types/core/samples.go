package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/utils"
	"time"
)

func (c *Core) Samples() []database.DbObject {
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

	return []database.DbObject{core}

}
