package failures

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

func Samples() {
	tx := DB().Begin()
	sg := new(sync.WaitGroup)

	createdAt := time.Now().Add(-1 * types.Month)

	for i := int64(1); i <= 4; i++ {
		sg.Add(1)

		log.Infoln(fmt.Sprintf("Adding %v Failure records to service", 730))

		go func() {
			defer sg.Done()
			for fi := 0.; fi <= float64(730); fi++ {
				createdAt = createdAt.Add(2 * time.Minute)
				failure := &Failure{
					Service:   i,
					Issue:     "testing right here",
					CreatedAt: createdAt.UTC(),
				}

				tx = tx.Create(&failure)
			}
		}()
	}
	sg.Wait()

	if err := tx.Commit().Error(); err != nil {
		log.Error(err)
	}

}
