package hits

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SampleHits = 99900.

func Samples() {
	tx := DB().Begin()
	sg := new(sync.WaitGroup)

	for i := int64(1); i <= 5; i++ {
		err := createHitsAt(tx, i, sg)
		if err != nil {
			log.Error(err)
		}
		tx = DB().Begin()
	}

}

func createHitsAt(db database.Database, serviceID int64, sg *sync.WaitGroup) error {
	createdAt := utils.Now().Add(-3 * types.Day)
	p := utils.NewPerlin(2, 2, 5, utils.Now().UnixNano())

	i := 0
	for hi := 0.; hi <= SampleHits; hi++ {
		latency := p.Noise1D(hi / 500)

		createdAt = createdAt.Add(30 * time.Second)

		hit := &Hit{
			Service:   serviceID,
			Latency:   latency,
			PingTime:  latency * 0.15,
			CreatedAt: createdAt,
		}

		db = db.Create(&hit)
		i++
		if createdAt.After(utils.Now()) {
			break
		}
	}

	return db.Commit().Error()
}
