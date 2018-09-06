// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifiers

import (
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testNotifier *Tester
	dir          string
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
	dir = utils.Directory
	utils.InitLogs()
}

func injectDatabase() {
	dbSession, _ := gorm.Open("sqlite3", dir+"/statup.db")
	Collections = dbSession.Table("communication").Model(&Notification{})
}

type Tester struct {
	*Notification
}

func TestInit(t *testing.T) {
	injectDatabase()
}

func TestAdd(t *testing.T) {
	testNotifier = &Tester{&Notification{
		Id:     999999,
		Method: "tester",
		Host:   "0.0.0.0",
		Form: []NotificationForm{{
			Id:          999999,
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
			DbField:     "Host",
		}}},
	}

	add(testNotifier)
}

func TestIsInDatabase(t *testing.T) {
	in := testNotifier.IsInDatabase()
	assert.False(t, in)
}

func TestInsertDatabase(t *testing.T) {
	newId, err := InsertDatabase(testNotifier.Notification)
	assert.Nil(t, err)
	assert.NotZero(t, newId)

	in := testNotifier.IsInDatabase()
	assert.True(t, in)
}

func TestSelectNotification(t *testing.T) {
	notifier, err := SelectNotification(999999)
	assert.Nil(t, err)
	assert.Equal(t, "tester", notifier.Method)
	assert.False(t, notifier.Enabled)
}

func TestNotification_Update(t *testing.T) {
	notifier, err := SelectNotification(999999)
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
	notifier, err := SelectNotification(999999)
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
