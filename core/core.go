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

package core

import (
	"errors"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"os"
	"time"
)

type PluginJSON types.PluginJSON
type PluginRepos types.PluginRepos

type Core struct {
	*types.Core
}

var (
	Configs   *DbConfig // Configs holds all of the config.yml and database info
	CoreApp   *Core     // CoreApp is a global variable that contains many elements
	SetupMode bool      // SetupMode will be true if Statup does not have a database connection
	VERSION   string    // VERSION is set on build automatically by setting a -ldflag
)

func init() {
	CoreApp = NewCore()
}

// NewCore return a new *core.Core struct
func NewCore() *Core {
	CoreApp = new(Core)
	CoreApp.Core = new(types.Core)
	CoreApp.Started = time.Now()
	return CoreApp
}

// ToCore will convert *core.Core to *types.Core
func (c *Core) ToCore() *types.Core {
	return c.Core
}

// InitApp will initialize Statup
func InitApp() {
	SelectCore()
	insertNotifierDB()
	CoreApp.SelectAllServices()
	checkServices()
	CoreApp.Notifications = notifier.Load()
	go DatabaseMaintence()
}

// insertNotifierDB inject the Statup database instance to the Notifier package
func insertNotifierDB() error {
	if DbSession == nil {
		err := Configs.Connect(false, utils.Directory)
		if err != nil {
			return errors.New("database connection has not been created")
		}
	}
	notifier.SetDB(DbSession)
	return nil
}

// UpdateCore will update the CoreApp variable inside of the 'core' table in database
func UpdateCore(c *Core) (*Core, error) {
	db := coreDB().Update(&c)
	return c, db.Error
}

// UsingAssets will return true if /assets folder is present
func (c Core) CurrentTime() string {
	t := time.Now().UTC()
	current := utils.Timezoner(t, c.Timezone)
	ansic := "Monday 03:04:05 PM"
	return current.Format(ansic)
}

// UsingAssets will return true if /assets folder is present
func (c Core) UsingAssets() bool {
	return source.UsingAssets(utils.Directory)
}

// SassVars opens the file /assets/scss/variables.scss to be edited in Theme
func (c Core) SassVars() string {
	if !source.UsingAssets(utils.Directory) {
		return ""
	}
	return source.OpenAsset(utils.Directory, "scss/variables.scss")
}

// BaseSASS is the base design , this opens the file /assets/scss/base.scss to be edited in Theme
func (c Core) BaseSASS() string {
	if !source.UsingAssets(utils.Directory) {
		return ""
	}
	return source.OpenAsset(utils.Directory, "scss/base.scss")
}

// MobileSASS is the -webkit responsive custom css designs. This opens the
// file /assets/scss/mobile.scss to be edited in Theme
func (c Core) MobileSASS() string {
	if !source.UsingAssets(utils.Directory) {
		return ""
	}
	return source.OpenAsset(utils.Directory, "scss/mobile.scss")
}

// AllOnline will be true if all services are online
func (c Core) AllOnline() bool {
	for _, s := range CoreApp.Services {
		if !s.Select().Online {
			return false
		}
	}
	return true
}

// SelectCore will return the CoreApp global variable and the settings/configs for Statup
func SelectCore() (*Core, error) {
	if DbSession == nil {
		return nil, errors.New("database has not been initiated yet.")
	}
	exists := DbSession.HasTable("core")
	if !exists {
		return nil, errors.New("core database has not been setup yet.")
	}
	db := coreDB().First(&CoreApp)
	if db.Error != nil {
		return nil, db.Error
	}
	CoreApp.DbConnection = Configs.DbConn
	CoreApp.Version = VERSION
	if os.Getenv("USE_CDN") == "true" {
		CoreApp.UseCdn = true
	}
	//store = sessions.NewCookieStore([]byte(core.ApiSecret))
	return CoreApp, db.Error
}

// ServiceOrder will reorder the services based on 'order_id' (Order)
type ServiceOrder []types.ServiceInterface

func (c ServiceOrder) Len() int           { return len(c) }
func (c ServiceOrder) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ServiceOrder) Less(i, j int) bool { return c[i].(*Service).Order < c[j].(*Service).Order }
