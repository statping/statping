package checkins

import (
	"github.com/statping/statping/utils"
	"time"
)

func Samples() error {
	log.Infoln("Inserting Sample Checkins...")
	checkin1 := &Checkin{
		Name:        "Demo Checkin 1",
		ServiceId:   1,
		Interval:    300,
		GracePeriod: 300,
		ApiKey:      "demoCheckin123",
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
	log.Infoln("Inserting Sample Checkins Hits...")
	checkTime := utils.Now().Add(-3 * time.Minute)

	for i := int64(1); i <= 2; i++ {
		checkHit := &CheckinHit{
			Checkin:   i,
			From:      "192.168.0.1",
			CreatedAt: checkTime.UTC(),
		}

		if err := checkHit.Create(); err != nil {
			return err
		}

		checkTime = checkTime.Add(1 * time.Minute)
	}

	return nil
}
