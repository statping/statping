package failures

import (
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/prometheus/common/log"
	"time"
)

func (f *Failure) Samples() []database.DbObject {

	createdAt := time.Now().Add(-1 * types.Month)

	for i := int64(1); i <= 5; i++ {
		log.Infoln(fmt.Sprintf("Adding %v Failure records to service", 5500))

		for fi := 1; fi <= 5500; fi++ {
			createdAt = createdAt.Add(2 * time.Minute)

			failure := &Failure{
				Service:   i,
				Issue:     "testing right here",
				CreatedAt: createdAt,
			}

			failure.Create()
		}
	}
	return nil
}
