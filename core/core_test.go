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
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCore   *Core
	testConfig *DbConfig
	dir        string
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	source.Assets()
}

func TestNewCore(t *testing.T) {
	testCore = NewCore()
	assert.NotNil(t, testCore)
	testCore.Name = "Tester"
}

func TestDbConfig_Save(t *testing.T) {
	testConfig = &DbConfig{&types.DbConfig{
		DbConn:   "sqlite",
		Project:  "Tester",
		Location: dir,
	}}
	err := testConfig.Save()
	assert.Nil(t, err)
}

func TestDbConnection(t *testing.T) {
	err := DbConnection(testConfig.DbConn, false, dir)
	assert.Nil(t, err)
}

func TestCreateDatabase(t *testing.T) {
	err := CreateDatabase()
	assert.Nil(t, err)
}

func TestInsertCore(t *testing.T) {
	err := InsertCore(testCore)
	assert.Nil(t, err)
}

func TestSelectCore(t *testing.T) {
	core, err := SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Tester", core.Name)
}

func TestSampleData(t *testing.T) {
	err := LoadSampleData()
	assert.Nil(t, err)
}

func TestSelectLastMigration(t *testing.T) {
	id, err := SelectLastMigration()
	assert.Nil(t, err)
	assert.NotZero(t, id)
}

func TestInsertNotifierDB(t *testing.T) {
	err := InsertNotifierDB()
	assert.Nil(t, err)
}

func TestExportStaticHTML(t *testing.T) {
	t.SkipNow()
	data := ExportIndexHTML()
	assert.Contains(t, data, "Statup  made with ❤️")
	assert.Contains(t, data, "</body>")
	assert.Contains(t, data, "</html>")
}
