// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var (
	dir          string
	db           *gorm.DB
	currentCount int
)

var TestService = &types.Service{
	Id:             1,
	Name:           "Interpol - All The Rage Back Home",
	Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
	ExpectedStatus: 200,
	Expected:       types.NewNullString("test example"),
	Interval:       30,
	Type:           "http",
	Method:         "GET",
	Timeout:        20,
	LastStatusCode: 404,
	Online:         true,
	LastResponse:   "<html>this is an example response</html>",
	CreatedAt:      utils.Now().Add(-24 * time.Hour),
}

var TestFailure = &types.Failure{
	Issue:     "testing",
	Service:   1,
	CreatedAt: utils.Now().Add(-12 * time.Hour),
}

var TestUser = &types.User{
	Username: "admin",
	Email:    "info@email.com",
}

var TestCore = &types.Core{
	Name: "testing notifiers",
}

func CountNotifiers() int {
	return len(notifier.AllCommunications)
}

func init() {
	dir = utils.Directory
	source.Assets()
	utils.InitLogs()
	injectDatabase()
}

func injectDatabase() {
	utils.DeleteFile(dir + "/notifiers.db")
	db, err := gorm.Open("sqlite3", dir+"/notifiers.db")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&notifier.Notification{})
	notifier.SetDB(db, float32(-8))
}
