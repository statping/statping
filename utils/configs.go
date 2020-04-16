package utils

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	Params *viper.Viper
)

func InitCLI() {
	Params = viper.New()
	setDefaults()
	Params.SetConfigName("config")
	Params.SetConfigType("yml")
	Params.AddConfigPath(".")
	err := Params.ReadInConfig()
	if err != nil {
		log.Debugf("config.yml Fatal error config file: %s", err)
	}

	Params.AddConfigPath(".")
	Params.SetConfigFile(".env")
	err = Params.ReadInConfig()
	if err != nil {
		log.Debugf(".env Fatal error config file: %s", err)
	}

	Params.AutomaticEnv()
	if err != nil {
		log.Debugf("No environment variables found: %s", err)
	}
}

func setDefaults() {
	defaultDir, err := os.Getwd()
	if err != nil {
		defaultDir = "."
	}
	Params.SetDefault("STATPING_DIR", defaultDir)
	Directory = Params.GetString("STATPING_DIR")
	Params.SetDefault("GO_ENV", "")
	Params.SetDefault("DISABLE_LOGS", false)
	Params.SetDefault("BASE_PATH", "")
	Params.SetDefault("MAX_OPEN_CONN", 25)
	Params.SetDefault("MAX_IDLE_CONN", 25)
	Params.SetDefault("MAX_LIFE_CONN", 5*time.Minute)
	Params.SetDefault("SAMPLE_DATA", true)
	Params.SetDefault("USE_CDN", false)
	Params.SetDefault("ALLOW_REPORTS", false)
	Params.SetDefault("AUTH_USERNAME", "")
	Params.SetDefault("AUTH_PASSWORD", "")
	Params.SetDefault("POSTGRES_SSLMODE", "disable")

	dbConn := Params.GetString("DB_CONN")
	dbInt := Params.GetInt("DB_PORT")
	if dbInt == 0 && dbConn != "sqlite" && dbConn != "sqlite3" {
		if dbConn == "postgres" {
			Params.SetDefault("DB_PORT", 5432)
		}
		if dbConn == "mysql" {
			Params.SetDefault("DB_PORT", 3306)
		}
	}
}
