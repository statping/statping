package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbConnection(t *testing.T) {
	err := CoreApp.Connect(configs, false, dir)
	assert.Nil(t, err)
}

func TestDropDatabase(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := CoreApp.DropDatabase()
	assert.Nil(t, err)
}

func TestSeedSchemaDatabase(t *testing.T) {
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
	err := InsertLargeSampleData()
	assert.Nil(t, err)
}

func TestReLoadDbConfig(t *testing.T) {
	err := CoreApp.Connect(configs, false, dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", CoreApp.config.DbConn)
}
