package hits

import (
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SampleHits = 9900.

func Samples() {
	tx := DB().Begin()
	sg := new(sync.WaitGroup)

	for i := int64(1); i <= 4; i++ {
		err := createHitsAt(tx, i, sg)
		if err != nil {
			log.Error(err)
		}
		tx = DB().Begin()
	}

}

func createHitsAt(db database.Database, serviceID int64, sg *sync.WaitGroup) error {
	createdAt := utils.Now().Add(-30 * types.Day)
	p := utils.NewPerlin(2, 2, 5, utils.Now().UnixNano())

	i := 0
	for hi := 0.; hi <= SampleHits; hi++ {
		latency := p.Noise1D(hi / 500)

		createdAt = createdAt.Add(10 * time.Minute)

		hit := &Hit{
			Service:   serviceID,
			Latency:   latency,
			PingTime:  latency * 0.15,
			CreatedAt: createdAt,
		}

		db = db.Create(&hit)
		fmt.Printf("Creating hit %d hit %d: %.2f %v\n", serviceID, hit.Id, latency, createdAt.String())
		i++
	}

	return db.Commit().Error()

}
