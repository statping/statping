package configs

import (
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func init() {
	utils.InitLogs()
}

func TestSQLiteConfig(t *testing.T) {
	sqlite := &DbConfig{
		DbConn: "sqlite",
		DbHost: "localhost",
		DbUser: "",
		DbPass: "",
		DbData: "",
		DbPort: 0,
	}

	err := Connect(sqlite, false)
	require.Nil(t, err)
}

func TestMySQLConfig(t *testing.T) {
	mysql := &DbConfig{
		DbConn: "mysql",
		DbHost: "localhost",
		DbUser: "root",
		DbPass: "password123",
		DbData: "statping",
		DbPort: 3306,
	}

	err := Connect(mysql, false)
	require.Nil(t, err)
}

func TestPostgresConfig(t *testing.T) {
	postgres := &DbConfig{
		DbConn: "postgres",
		DbHost: "localhost",
		DbUser: "root",
		DbPass: "password123",
		DbData: "statping",
		DbPort: 5432,
	}

	err := Connect(postgres, false)
	require.Nil(t, err)
}
