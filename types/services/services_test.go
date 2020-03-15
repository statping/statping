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

func TestInit(t *testing.T) {
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
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Example Service", item.Name)
	assert.NotZero(t, item.LastOnline)
	assert.NotZero(t, item.LastOffline)
	assert.NotZero(t, item.LastCheck)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestService_Checkins(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Len(t, item.Checkins(), 1)
}

func TestService_AllHits(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Len(t, item.AllHits().List(), 3)
	assert.Equal(t, 3, item.AllHits().Count())
}

func TestService_AllFailures(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Len(t, item.AllFailures().List(), 2)
	assert.Equal(t, 2, item.AllFailures().Count())
}

func TestService_FirstHit(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	hit := item.FirstHit()
	assert.Equal(t, int64(1), hit.Id)
}

func TestService_LastHit(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	hit := item.AllHits().Last()
	assert.Equal(t, int64(3), hit.Id)
}

func TestService_LastFailure(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	fail := item.AllFailures().Last()
	assert.Equal(t, int64(2), fail.Id)
}

func TestService_Duration(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, float64(30), item.Duration().Seconds())
}

func TestService_CountHits(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	count := item.AllHits().Count()
	assert.NotZero(t, count)
}

func TestService_AvgTime(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)

	assert.Equal(t, int64(123456), item.AvgTime())
}

func TestService_HitsSince(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)

	count := item.HitsSince(utils.Now().Add(-30 * time.Second))
	assert.Equal(t, 1, count.Count())

	count = item.HitsSince(utils.Now().Add(-180 * time.Second))
	assert.Equal(t, 3, count.Count())
}

func TestService_IsRunning(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.False(t, item.IsRunning())
}

func TestService_OnlineDaysPercent(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)

	amount := item.OnlineDaysPercent(1)

	assert.Equal(t, float32(33.33), amount)
}

func TestService_Downtime(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	amount := item.Downtime().Seconds()
	assert.Equal(t, "25", fmt.Sprintf("%0.f", amount))
}

func TestService_FailuresSince(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)

	count := item.FailuresSince(utils.Now().Add(-6 * time.Second))
	assert.Equal(t, 1, count.Count())

	count = item.FailuresSince(utils.Now().Add(-180 * time.Second))
	assert.Equal(t, 2, count.Count())
}

func TestCreate(t *testing.T) {
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
}

func TestUpdate(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	item.Name = "Updated Service"
	item.Order = 1
	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, int64(1), item.Id)
	assert.Equal(t, "Updated Service", item.Name)
}

func TestAllInOrder(t *testing.T) {
	inOrder := AllInOrder()
	assert.Len(t, inOrder, 2)
	assert.Equal(t, "Updated Service", inOrder[0].Name)
	assert.Equal(t, "Example Service 2", inOrder[1].Name)
}

func TestDelete(t *testing.T) {
	all := All()
	assert.Len(t, all, 2)

	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, int64(1), item.Id)

	err = item.Delete()
	require.Nil(t, err)

	all = All()
	assert.Len(t, all, 1)
}

func TestService_CheckService(t *testing.T) {
	item, err := Find(2)
	require.Nil(t, err)

	hitsCount := item.AllHits().Count()
	failsCount := item.AllFailures().Count()

	assert.Equal(t, 3, hitsCount)
	assert.Equal(t, 2, failsCount)

	item.CheckService(true)

	assert.Equal(t, 4, hitsCount)
	assert.Equal(t, 2, failsCount)
}

func TestClose(t *testing.T) {
	assert.Nil(t, db.Close())
}
