package checkins

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
	"time"
)

func (c *Checkin) Samples() []database.DbObject {
	checkin1 := &Checkin{
		Name:        "Example Checkin 1",
		ServiceId:   1,
		Interval:    300,
		GracePeriod: 300,
		ApiKey:      utils.RandomString(7),
	}

	checkin2 := &Checkin{
		Name:        "Example Checkin 2",
		ServiceId:   2,
		Interval:    900,
		GracePeriod: 300,
		ApiKey:      utils.RandomString(7),
	}

	return []database.DbObject{checkin1, checkin2}
}

func (c *CheckinHit) Samples() []database.DbObject {
	checkTime := time.Now().UTC().Add(-24 * time.Hour)

	var hits []database.DbObject

	for i := int64(1); i <= 2; i++ {
		checkHit := &CheckinHit{
			Checkin:   i,
			From:      "192.168.0.1",
			CreatedAt: checkTime.UTC(),
		}

		hits = append(hits, checkHit)

		checkTime = checkTime.Add(10 * time.Minute)
	}

	return hits
}
