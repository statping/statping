package hits

import (
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"sync"
	"time"
)

var SampleHits = 9900.

func Samples() {
	tx := DB().Begin()
	sg := new(sync.WaitGroup)

	createdAt := time.Now().Add(-1 * types.Month)
	for i := int64(1); i <= 4; i++ {
		sg.Add(1)

		p := utils.NewPerlin(2, 2, 5, time.Now().UnixNano())

		go func() {
			defer sg.Done()
			for hi := 0.; hi <= SampleHits; hi++ {
				latency := p.Noise1D(hi / 500)
				createdAt = createdAt.Add(1 * time.Minute)
				hit := &Hit{
					Service:   i,
					CreatedAt: createdAt.UTC(),
					Latency:   latency,
				}
				tx = tx.Create(&hit)
			}
		}()
	}
	sg.Wait()

	if err := tx.Commit().Error(); err != nil {
		log.Error(err)
	}
}
