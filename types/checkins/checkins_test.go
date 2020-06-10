package checkins

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testCheckin = &Checkin{
	ServiceId:   1,
	Name:        "Test Checkin",
	Interval:    60,
	GracePeriod: 10,
	ApiKey:      "tHiSiSaTeStXXX",
	CreatedAt:   utils.Now(),
	UpdatedAt:   utils.Now(),
	LastHitTime: utils.Now().Add(-15 * time.Second),
}

var testCheckinHits = []*CheckinHit{{
	Checkin:   1,
	From:      "0.0.0.0",
	CreatedAt: utils.Now().Add(-30 * time.Second),
}, {
	Checkin:   2,
	From:      "0.0.0.0",
	CreatedAt: utils.Now().Add(-180 * time.Second),
}}

var testApiKey string

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	SetDB(db)
	db.AutoMigrate(&Checkin{}, &CheckinHit{}, &failures.Failure{})
	db.Create(&testCheckin)
	for _, v := range testCheckinHits {
		db.Create(&v)
	}
	assert.True(t, db.HasTable(&Checkin{}))
	assert.True(t, db.HasTable(&CheckinHit{}))
	assert.True(t, db.HasTable(&failures.Failure{}))

	t.Run("Test Checkin", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)
		assert.Equal(t, "Test Checkin", item.Name)
		assert.NotEmpty(t, item.ApiKey)
		testApiKey = item.ApiKey
	})

	t.Run("Test FindByAPI", func(t *testing.T) {
		item, err := FindByAPI(testApiKey)
		require.Nil(t, err)
		assert.Equal(t, "Test Checkin", item.Name)
	})

	t.Run("Test All", func(t *testing.T) {
		items := All()
		assert.Len(t, items, 1)
	})

	t.Run("Test Create", func(t *testing.T) {
		example := &Checkin{
			Name: "Example 2",
		}
		err := example.Create()
		example.Close()
		require.Nil(t, err)
		assert.NotZero(t, example.Id)
		assert.Equal(t, "Example 2", example.Name)
		assert.NotZero(t, example.CreatedAt)
		assert.NotEmpty(t, example.ApiKey)
	})

	t.Run("Test Update", func(t *testing.T) {
		i, err := Find(1)
		require.Nil(t, err)
		i.Name = "Updated"

		err = i.Update()
		require.Nil(t, err)
		assert.Equal(t, "Updated", i.Name)
	})

	t.Run("Test Expected Time", func(t *testing.T) {
		item, err := Find(1)
		require.Nil(t, err)

		expected := item.Expected()
		assert.GreaterOrEqual(t, expected.Seconds(), float64(29))
	})

	t.Run("Test Delete", func(t *testing.T) {
		all := All()
		assert.Len(t, all, 2)

		item, err := Find(2)
		require.Nil(t, err)

		err = item.Delete()
		require.Nil(t, err)

		all = All()
		assert.Len(t, all, 1)
	})

	t.Run("Test Checkin", func(t *testing.T) {
		assert.Nil(t, db.Close())
		assert.Nil(t, dbHits.Close())
	})

}
