package hits

import (
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

var SampleHits = 99900.

func Samples() error {
	tx := db.Begin()
	sg := new(sync.WaitGroup)

	for i := int64(1); i <= 5; i++ {
		err := createHitsAt(tx, i, sg)
		if err != nil {
			log.Error(err)
		}
		tx = db.Begin()
	}

	return tx.Error()
}

func createHitsAt(db database.Database, serviceID int64, sg *sync.WaitGroup) error {
	log.Infoln(fmt.Sprintf("Adding sample hit records to service #%d", serviceID))

	createdAt := utils.Now().Add(-3 * types.Day)
	p := utils.NewPerlin(2, 2, 5, utils.Now().UnixNano())

	i := 0
	for hi := 0.; hi <= SampleHits; hi++ {
		latency := p.Noise1D(hi / 500)

		createdAt = createdAt.Add(30 * time.Second)

		hit := &Hit{
			Service:   serviceID,
			Latency:   int64(latency * 10000000),
			PingTime:  int64(latency * 5000000),
			CreatedAt: createdAt,
		}

		db = db.Create(&hit)
		if err := db.Error(); err != nil {
			return err
		}

		i++
		if createdAt.After(utils.Now()) {
			break
		}
	}

	return db.Commit().Error()
}
