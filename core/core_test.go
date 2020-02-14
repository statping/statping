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

package core

import (
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	dir       string
	skipNewDb bool
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	source.Assets()
	skipNewDb = false
	SampleHits = 480
}

func TestNewCore(t *testing.T) {
	err := TmpRecords("core.db")
	require.Nil(t, err)
	require.NotNil(t, CoreApp)
}

func TestDbConfig_Save(t *testing.T) {
	t.SkipNow()
	//if skipNewDb {
	//	t.SkipNow()
	//}
	//var err error
	//Configs = &DbConfig{
	//	DbConn:   "sqlite",
	//	Project:  "Tester",
	//	Location: dir,
	//}
	//Configs, err = Configs.Save()
	//assert.Nil(t, err)
	//assert.Equal(t, "sqlite", Configs.DbConn)
	//assert.NotEmpty(t, Configs.ApiKey)
	//assert.NotEmpty(t, Configs.ApiSecret)
}

func TestLoadDbConfig(t *testing.T) {
	Configs, err := LoadConfigFile(dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", Configs.DbConn)
}

func TestDbConnection(t *testing.T) {
	err := CoreApp.Connect(false, dir)
	assert.Nil(t, err)
}

func TestDropDatabase(t *testing.T) {
	t.SkipNow()
	if skipNewDb {
		t.SkipNow()
	}
	err := CoreApp.DropDatabase()
	assert.Nil(t, err)
}

func TestSeedSchemaDatabase(t *testing.T) {
	t.SkipNow()
	if skipNewDb {
		t.SkipNow()
	}
	err := CoreApp.CreateDatabase()
	assert.Nil(t, err)
}

func TestMigrateDatabase(t *testing.T) {
	t.SkipNow()
	err := CoreApp.MigrateDatabase()
	assert.Nil(t, err)
}

func TestSeedDatabase(t *testing.T) {
	t.SkipNow()
	err := InsertLargeSampleData()
	assert.Nil(t, err)
}

func TestReLoadDbConfig(t *testing.T) {
	err := CoreApp.Connect(false, dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", CoreApp.Config.DbConn)
}

func TestSelectCore(t *testing.T) {
	core, err := SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Statping Sample Data", core.Name)
}

func TestInsertNotifierDB(t *testing.T) {
	t.SkipNow()
	if skipNewDb {
		t.SkipNow()
	}
	err := InsertNotifierDB()
	assert.Nil(t, err)
}

func TestEnvToConfig(t *testing.T) {
	os.Setenv("DB_CONN", "sqlite")
	os.Setenv("DB_USER", "")
	os.Setenv("DB_PASS", "")
	os.Setenv("DB_DATABASE", "")
	os.Setenv("NAME", "Testing")
	os.Setenv("DOMAIN", "http://localhost:8080")
	os.Setenv("DESCRIPTION", "Testing Statping")
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASS", "admin123")
	os.Setenv("VERBOSE", "1")
	config, err := EnvToConfig()
	assert.Nil(t, err)
	assert.Equal(t, config.DbConn, "sqlite")
	assert.Equal(t, config.Domain, "http://localhost:8080")
	assert.Equal(t, config.Description, "Testing Statping")
	assert.Equal(t, config.Username, "admin")
	assert.Equal(t, config.Password, "admin123")
}

func TestGetLocalIP(t *testing.T) {
	ip := GetLocalIP()
	assert.Contains(t, ip, "http://")
}
