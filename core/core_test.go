package core

import (
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"os"
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
	testConfig = &DbConfig{
		DbConn:   "sqlite",
		Project:  "Tester",
		Location: dir,
	}
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

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, testCore.UsingAssets())
}

func TestHasAssets(t *testing.T) {
	assert.False(t, HasAssets(dir))
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, HasAssets(dir))
}

func TestCompileSASS(t *testing.T) {
	t.SkipNow()
	os.Setenv("SASS", "sass")
	os.Setenv("CMD_FILE", dir+"/cmd.sh")
	assert.Nil(t, CompileSASS(dir))
	assert.True(t, HasAssets(dir))
}

func TestDeleteAssets(t *testing.T) {
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, HasAssets(dir))
}

func TestInsertNotifierDB(t *testing.T) {
	err := InsertNotifierDB()
	assert.Nil(t, err)
}
