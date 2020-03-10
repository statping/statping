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
	From:      "0.0.0.0.0",
	CreatedAt: utils.Now().Add(-30 * time.Second),
}, {
	Checkin:   2,
	From:      "0.0.0.0",
	CreatedAt: utils.Now().Add(-180 * time.Second),
}}

var testApiKey string

func TestInit(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&Checkin{}, &CheckinHit{}, &failures.Failure{})
	db.Create(&testCheckin)
	for _, v := range testCheckinHits {
		db.Create(&v)
	}
	assert.True(t, db.HasTable(&Checkin{}))
	assert.True(t, db.HasTable(&CheckinHit{}))
	assert.True(t, db.HasTable(&failures.Failure{}))
	SetDB(db)
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Test Checkin", item.Name)
	assert.NotEmpty(t, item.ApiKey)
	testApiKey = item.ApiKey
}

func TestFindByAPI(t *testing.T) {
	item, err := FindByAPI(testApiKey)
	require.Nil(t, err)
	assert.Equal(t, "Test Checkin", item.Name)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestCreate(t *testing.T) {
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
}

func TestUpdate(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	item.Name = "Updated"

	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, "Updated", item.Name)
	item.Close()
}

func TestCheckin_Expected(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)

	expected := item.Expected()
	assert.GreaterOrEqual(t, expected.Seconds(), float64(29))
}

func TestDelete(t *testing.T) {
	all := All()
	assert.Len(t, all, 2)

	item, err := Find(2)
	require.Nil(t, err)

	err = item.Delete()
	require.Nil(t, err)

	all = All()
	assert.Len(t, all, 1)
}

func TestClose(t *testing.T) {
	assert.Nil(t, db.Close())
}
