package services

import (
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var example = &Service{
	Name:           "Example Service",
	Domain:         "https://statping.com",
	ExpectedStatus: 200,
	Interval:       30,
	Type:           "http",
	Method:         "GET",
	Timeout:        5,
	Order:          1,
	VerifySSL:      null.NewNullBool(true),
	Public:         null.NewNullBool(true),
	GroupId:        1,
	Permalink:      null.NewNullString("statping"),
	LastCheck:      utils.Now().Add(-5 * time.Second),
	LastOffline:    utils.Now().Add(-5 * time.Second),
	LastOnline:     utils.Now().Add(-60 * time.Second),
}

var hit1 = &hits.Hit{
	Service:   1,
	Latency:   123456,
	PingTime:  123456,
	CreatedAt: utils.Now().Add(-120 * time.Second),
}

var hit2 = &hits.Hit{
	Service:   1,
	Latency:   123456,
	PingTime:  123456,
	CreatedAt: utils.Now().Add(-60 * time.Second),
}

var hit3 = &hits.Hit{
	Service:   1,
	Latency:   123456,
	PingTime:  123456,
	CreatedAt: utils.Now().Add(-30 * time.Second),
}

var exmapleCheckin = &checkins.Checkin{
	ServiceId:   1,
	Name:        "Example Checkin",
	Interval:    60,
	GracePeriod: 30,
	ApiKey:      "wdededede",
}

var fail1 = &failures.Failure{
	Issue:     "example not found",
	ErrorCode: 404,
	Service:   1,
	PingTime:  123456,
	CreatedAt: utils.Now().Add(-160 * time.Second),
}

var fail2 = &failures.Failure{
	Issue:     "example 2 not found",
	ErrorCode: 500,
	Service:   1,
	PingTime:  123456,
	CreatedAt: utils.Now().Add(-5 * time.Second),
}

func TestServices(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&Service{}, &hits.Hit{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &failures.Failure{})
	db.Create(&example)
	db.Create(&hit1)
	db.Create(&hit2)
	db.Create(&hit3)
	db.Create(&exmapleCheckin)
	db.Create(&fail1)
	db.Create(&fail2)
	checkins.SetDB(db)
	failures.SetDB(db)
	hits.SetDB(db)
	SetDB(db)

	t.Run("Test Find service", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Equal(t, "Example Service", item.Name)
		assert.NotZero(t, item.LastOnline)
		assert.NotZero(t, item.LastOffline)
		assert.NotZero(t, item.LastCheck)
	})

	t.Run("Test All", func(t *testing.T) {
		items := All()
		assert.Len(t, items, 1)
	})

	t.Run("Test Checkins", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Len(t, item.Checkins(), 1)
	})

	t.Run("Test All Hits", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Len(t, item.AllHits().List(), 3)
		assert.Equal(t, 3, item.AllHits().Count())
	})

	t.Run("Test All Failures", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Len(t, item.AllFailures().List(), 2)
		assert.Equal(t, 2, item.AllFailures().Count())
	})

	t.Run("Test First Hit", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		hit := item.FirstHit()
		assert.Equal(t, int64(1), hit.Id)
	})

	t.Run("Test Last Hit", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		hit := item.AllHits().Last()
		assert.Equal(t, int64(3), hit.Id)
	})

	t.Run("Test Last Failure", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		fail := item.AllFailures().Last()
		assert.Equal(t, int64(2), fail.Id)
	})

	t.Run("Test Duration", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Equal(t, float64(30), item.Duration().Seconds())
	})

	t.Run("Test Count Hits", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		count := item.AllHits().Count()
		assert.NotZero(t, count)
	})

	t.Run("Test Average Time", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)

		assert.Equal(t, int64(123456), item.AvgTime())
	})

	t.Run("Test Hits Since", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)

		count := item.HitsSince(utils.Now().Add(-30 * time.Second))
		assert.Equal(t, 1, count.Count())

		count = item.HitsSince(utils.Now().Add(-180 * time.Second))
		assert.Equal(t, 3, count.Count())
	})

	t.Run("Test Service Running", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.False(t, item.IsRunning())
	})

	t.Run("Test Online Percent", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)

		amount := item.OnlineDaysPercent(1)

		assert.Equal(t, float32(33.33), amount)
	})

	t.Run("Test Downtime", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		amount := item.Downtime().Seconds()
		assert.Equal(t, "25", fmt.Sprintf("%0.f", amount))
	})

	t.Run("Test Failures Since", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)

		count := item.FailuresSince(utils.Now().Add(-6 * time.Second))
		assert.Equal(t, 1, count.Count())

		count = item.FailuresSince(utils.Now().Add(-180 * time.Second))
		assert.Equal(t, 2, count.Count())
	})

	t.Run("Test Create", func(t *testing.T) {
		example := &Service{
			Name:           "Example Service 2",
			Domain:         "https://slack.statping.com",
			ExpectedStatus: 200,
			Interval:       10,
			Type:           "http",
			Method:         "GET",
			Timeout:        5,
			Order:          3,
			VerifySSL:      null.NewNullBool(true),
			Public:         null.NewNullBool(false),
			GroupId:        1,
			Permalink:      null.NewNullString("statping2"),
		}
		err := example.Create()
		require.Nil(t, err)
		assert.NotZero(t, example.Id)
		assert.Equal(t, "Example Service 2", example.Name)
		assert.False(t, example.Public.Bool)
		assert.NotZero(t, example.CreatedAt)
		assert.Equal(t, int64(2), example.Id)
		assert.Len(t, allServices, 2)
	})

	t.Run("Test Update Service", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		item.Name = "Updated Service"
		item.Order = 1
		err = item.Update()
		require.Nil(t, err)
		assert.Equal(t, int64(1), item.Id)
		assert.Equal(t, "Updated Service", item.Name)
	})

	t.Run("Test In Order", func(t *testing.T) {
		inOrder := AllInOrder()
		assert.Len(t, inOrder, 2)
		assert.Equal(t, "Updated Service", inOrder[0].Name)
		assert.Equal(t, "Example Service 2", inOrder[1].Name)
	})

	t.Run("Test Delete", func(t *testing.T) {
		all := All()
		assert.Len(t, all, 2)

		item, err := Find(1)
		require.Nil(t, err)
		assert.Equal(t, int64(1), item.Id)

		err = item.Delete()
		require.Nil(t, err)

		all = All()
		assert.Len(t, all, 1)
	})

	t.Run("Test Close", func(t *testing.T) {
		assert.Nil(t, db.Close())
	})

}
