package hits

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"time"
)

func (u *Hit) Samples() []database.DbObject {
	createdAt := time.Now().Add(-1 * types.Month)
	var hits []database.DbObject

	for i := int64(1); i <= 5; i++ {
		p := utils.NewPerlin(2, 2, 5, time.Now().UnixNano())

		for hi := 1; hi <= 5500; hi++ {
			latency := p.Noise1D(float64(hi / 10))
			createdAt = createdAt.Add(1 * time.Minute)
			hit := &Hit{
				Service:   i,
				CreatedAt: createdAt.UTC(),
				Latency:   latency,
			}
			hits = append(hits, hit)
		}
	}

	return hits
}
