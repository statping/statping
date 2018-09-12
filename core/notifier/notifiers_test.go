// +build bypass

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

package notifier

import (
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	dir        string
	EXAMPLE_ID = "example"
)

func init() {
	dir = utils.Directory
}

func injectDatabase() {
	db, _ = gorm.Open("sqlite3", dir+"/statup.db")
}

func TestLoad(t *testing.T) {
	source.Assets()
	utils.InitLogs()
	injectDatabase()
	AllCommunications = Load()
	assert.Len(t, AllCommunications, 1)
}

func TestIsInDatabase(t *testing.T) {
	in := example.isInDatabase()
	assert.True(t, in)
}

func TestInsertDatabase(t *testing.T) {
	_, err := insertDatabase(example.Notification)
	assert.Nil(t, err)
	assert.NotZero(t, example.Id)

	in := example.isInDatabase()
	assert.True(t, in)
}

func TestSelectNotification(t *testing.T) {
	notifier, err := SelectNotification(EXAMPLE_ID)
	assert.Nil(t, err)
	assert.Equal(t, "example", notifier.Method)
	assert.False(t, notifier.Enabled)
}

func TestNotification_Update(t *testing.T) {
	notifier, err := SelectNotification(EXAMPLE_ID)
	assert.Nil(t, err)
	notifier.Host = "new host here"
	notifier.Enabled = true
	updated, err := notifier.Update()
	assert.Nil(t, err)
	selected, err := SelectNotification(updated.Method)
	assert.Nil(t, err)
	assert.Equal(t, "new host here", selected.GetValue("host"))
	assert.True(t, selected.Enabled)
}

func TestNotification_GetValue(t *testing.T) {
	notifier, err := SelectNotification(EXAMPLE_ID)
	assert.Nil(t, err)
	val := notifier.GetValue("Host")
	assert.Equal(t, "http://exmaplehost.com", val)
}

//func TestRun(t *testing.T) {
//	err := example.Run()
//	assert.Equal(t, "running", err.Error())
//}

//func TestTestIt(t *testing.T) {
//	err := example.Test()
//	assert.Equal(t, "testing", err.Error())
//}

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
	f := &types.Failure{
		Issue: "testing",
	}
	OnFailure(s, f)
}
