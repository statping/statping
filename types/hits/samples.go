package hits

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	"time"
)

var SampleHits = 99900.

func Samples() error {
	log.Infoln("Inserting Sample Service Hits...")
	for i := int64(1); i <= 5; i++ {
		records := createHitsAt(i)
		if err := gormbulk.BulkInsert(db.GormDB(), records, db.ChunkSize()); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func createHitsAt(serviceID int64) []interface{} {
	log.Infoln(fmt.Sprintf("Adding Sample records to service #%d...", serviceID))

	createdAt := utils.Now().Add(-3 * types.Day)
	p := utils.NewPerlin(2, 2, 5, utils.Now().UnixNano())

	var records []interface{}
	for hi := 0.; hi <= SampleHits; hi++ {
		latency := p.Noise1D(hi / 500)

		hit := &Hit{
			Service:   serviceID,
			Latency:   int64(latency * 10000000),
			PingTime:  int64(latency * 5000000),
			CreatedAt: createdAt,
		}

		records = append(records, hit)

		if createdAt.After(utils.Now()) {
			break
		}
		createdAt = createdAt.Add(30 * time.Second)
	}

	return records
}
