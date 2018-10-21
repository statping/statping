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
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
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
}

func TestNewCore(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	utils.DeleteFile(dir + "/config.yml")
	utils.DeleteFile(dir + "/statup.db")
	CoreApp = NewCore()
	assert.NotNil(t, CoreApp)
	CoreApp.Name = "Tester"
}

func TestDbConfig_Save(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	var err error
	Configs = &DbConfig{
		DbConn:   "sqlite",
		Project:  "Tester",
		Location: dir,
	}
	Configs, err = Configs.Save()
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", Configs.DbConn)
	assert.NotEmpty(t, Configs.ApiKey)
	assert.NotEmpty(t, Configs.ApiSecret)
}

func TestLoadDbConfig(t *testing.T) {
	Configs, err := LoadConfigFile(dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", Configs.DbConn)
}

func TestDbConnection(t *testing.T) {
	err := Configs.Connect(false, dir)
	assert.Nil(t, err)
}

func TestDropDatabase(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := Configs.DropDatabase()
	assert.Nil(t, err)
}

func TestSeedSchemaDatabase(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := Configs.CreateDatabase()
	assert.Nil(t, err)
}

func TestMigrateDatabase(t *testing.T) {
	err := Configs.MigrateDatabase()
	assert.Nil(t, err)
}

func TestSeedDatabase(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := InsertLargeSampleData()
	assert.Nil(t, err)
}

func TestReLoadDbConfig(t *testing.T) {
	err := Configs.Connect(false, dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", Configs.DbConn)
}

func TestSelectCore(t *testing.T) {
	core, err := SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Statup Sample Data", core.Name)
}

func TestInsertNotifierDB(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := InsertNotifierDB()
	assert.Nil(t, err)
}

func TestExportStaticHTML(t *testing.T) {
	t.SkipNow()
	data := ExportIndexHTML()
	assert.Contains(t, data, "Statup  made with️")
	assert.Contains(t, data, "</body>")
	assert.Contains(t, data, "</html>")
}
