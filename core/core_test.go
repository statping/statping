package core

import (
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testCore     *Core
	testConfig   *DbConfig
	testDatabase string
)

func init() {
	testDatabase = os.Getenv("GOPATH")
	testDatabase += "/src/github.com/hunterlong/statup/"

	utils.InitLogs()
	RenderBoxes()
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
		Location: testDatabase,
	}
	err := testConfig.Save()
	assert.Nil(t, err)
}

func TestDbConnection(t *testing.T) {
	err := DbConnection(testConfig.DbConn, false, testDatabase)
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
	assert.False(t, HasAssets())
}

func TestInsertNotifierDB(t *testing.T) {
	err := InsertNotifierDB()
	assert.Nil(t, err)
}
