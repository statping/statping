package configs

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDbConfig_Save(t *testing.T) {
	config := &types.DbConfig{
		DbConn:   "sqlite",
		Project:  "Tester",
		Location: dir,
	}

	err := SaveConfig(config)
	require.Nil(t, err)
	assert.Equal(t, "sqlite", config.DbConn)
	assert.NotEmpty(t, config.ApiKey)
	assert.NotEmpty(t, config.ApiSecret)
}

func TestLoadDbConfig(t *testing.T) {
	Configs, err := LoadConfigFile(dir)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", Configs.DbConn)

	configs = Configs
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
