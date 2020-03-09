package checkins

import (
	"github.com/hunterlong/statping/utils"
	"time"
)

func Samples() error {
	checkin1 := &Checkin{
		Name:        "Example Checkin 1",
		ServiceId:   1,
		Interval:    300,
		GracePeriod: 300,
		ApiKey:      utils.RandomString(7),
	}
	if err := checkin1.Create(); err != nil {
		return err
	}

	checkin2 := &Checkin{
		Name:        "Example Checkin 2",
		ServiceId:   2,
		Interval:    900,
		GracePeriod: 300,
		ApiKey:      utils.RandomString(7),
	}
	if err := checkin2.Create(); err != nil {
		return err
	}
	return nil
}

func SamplesChkHits() error {
	checkTime := time.Now().UTC().Add(-24 * time.Hour)

	for i := int64(1); i <= 2; i++ {
		checkHit := &CheckinHit{
			Checkin:   i,
			From:      "192.168.0.1",
			CreatedAt: checkTime.UTC(),
		}

		if err := checkHit.Create(); err != nil {
			return err
		}

		checkTime = checkTime.Add(10 * time.Minute)
	}

	return nil
}
