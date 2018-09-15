package notifiers

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	dir          string
	db           *gorm.DB
	currentCount int
)

var TestService = &types.Service{
	Name:           "Interpol - All The Rage Back Home",
	Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
	ExpectedStatus: 200,
	Interval:       30,
	Type:           "http",
	Method:         "GET",
	Timeout:        20,
}

var TestFailure = &types.Failure{
	Issue: "testing",
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
	utils.DeleteFile(dir + "/statup.db")
	db, err := gorm.Open("sqlite3", dir+"/statup.db")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&notifier.Notification{})
	notifier.SetDB(db)
}
