package notifiers

import (
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"upper.io/db.v3/sqlite"
)

var (
	testNotifier *Tester
	testDatabase string
)

//
//
//

func (n *Tester) Init() error {
	return errors.New("im just testing")
}

func (n *Tester) Install() error {
	return errors.New("installing")
}

func (n *Tester) Run() error {
	return errors.New("running")
}

func (n *Tester) Select() *Notification {
	return n.Notification
}

func (n *Tester) OnSuccess(s *types.Service) error {
	return errors.New(s.Name)
}

func (n *Tester) OnFailure(s *types.Service) error {
	return errors.New(s.Name)
}

func (n *Tester) Test() error {
	return errors.New("testing")
}

func init() {
	testDatabase = os.Getenv("GOPATH")
	testDatabase += "/src/github.com/hunterlong/statup/"

	utils.InitLogs()
}

func injectDatabase() {
	sqliteDb := sqlite.ConnectionURL{
		Database: testDatabase + "statup.db",
	}
	dbSession, _ := sqlite.Open(sqliteDb)
	Collections = dbSession.Collection("communication")
}

type Tester struct {
	*Notification
}

func TestInit(t *testing.T) {
	injectDatabase()
}

func TestAdd(t *testing.T) {
	testNotifier = &Tester{&Notification{
		Id:     1,
		Method: "tester",
		Host:   "0.0.0.0",
		Form: []NotificationForm{{
			Id:          1,
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
			DbField:     "Host",
		}}},
	}

	add(testNotifier)
}

func TestIsInDatabase(t *testing.T) {
	in, err := testNotifier.IsInDatabase()
	assert.Nil(t, err)
	assert.False(t, in)
}

func TestInsertDatabase(t *testing.T) {
	newId, err := InsertDatabase(testNotifier.Notification)
	assert.Nil(t, err)
	assert.NotZero(t, newId)

	in, err := testNotifier.IsInDatabase()
	assert.Nil(t, err)
	assert.True(t, in)
}

func TestSelectNotification(t *testing.T) {
	notifier, err := SelectNotification(1)
	assert.Nil(t, err)
	assert.Equal(t, "tester", notifier.Method)
	assert.False(t, notifier.Enabled)
}

func TestNotification_Update(t *testing.T) {
	notifier, err := SelectNotification(1)
	assert.Nil(t, err)
	notifier.Method = "updatedName"
	notifier.Enabled = true
	updated, err := notifier.Update()
	assert.Nil(t, err)
	selected, err := SelectNotification(updated.Id)
	assert.Nil(t, err)
	assert.Equal(t, "updatedName", selected.Method)
	assert.True(t, selected.Enabled)
}

func TestNotification_GetValue(t *testing.T) {
	notifier, err := SelectNotification(1)
	assert.Nil(t, err)
	val := notifier.GetValue("Host")
	assert.Equal(t, "0.0.0.0", val)
}

func TestRun(t *testing.T) {
	err := testNotifier.Run()
	assert.Equal(t, "running", err.Error())
}

func TestTestIt(t *testing.T) {
	err := testNotifier.Test()
	assert.Equal(t, "testing", err.Error())
}

func TestOnSuccess(t *testing.T) {
	s := &types.Service{
		Name:           "Interpol - All The Rage Back Home",
		Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	}
	OnSuccess(s)
}

func TestOnFailure(t *testing.T) {
	s := &types.Service{
		Name:           "Interpol - All The Rage Back Home",
		Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	}
	OnFailure(s)
}
