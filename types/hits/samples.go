package hits

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
	"time"
)

var SampleHits = 99900.

func Samples() error {
	for i := int64(1); i <= 5; i++ {
		tx := db.Begin()
		tx = createHitsAt(tx, i)

		if err := tx.Commit().Error(); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func createHitsAt(db database.Database, serviceID int64) database.Database {
	log.Infoln(fmt.Sprintf("Adding Sample records to service #%d", serviceID))

	createdAt := utils.Now().Add(-3 * types.Day)
	p := utils.NewPerlin(2, 2, 5, utils.Now().UnixNano())

	for hi := 0.; hi <= SampleHits; hi++ {
		latency := p.Noise1D(hi / 500)

		hit := &Hit{
			Service:   serviceID,
			Latency:   int64(latency * 10000000),
			PingTime:  int64(latency * 5000000),
			CreatedAt: createdAt,
		}

		db = db.Create(&hit)

		if createdAt.After(utils.Now()) {
			break
		}
		createdAt = createdAt.Add(30 * time.Second)
	}

	return db
}
