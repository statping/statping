package utils

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	Params *viper.Viper
)

func InitCLI() {
	Params = viper.New()
	Params.AutomaticEnv()
	Directory = Params.GetString("STATPING_DIR")
	//Params.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	setDefaults()
	Params.SetConfigName("config")
	Params.SetConfigType("yml")
	Params.AddConfigPath(Directory)

	Params.ReadInConfig()

	Params.AddConfigPath(Directory)
	Params.SetConfigFile(".env")
	Params.ReadInConfig()

	Params.Set("VERSION", version)
}

func setDefaults() {
	if Directory == "" {
		defaultDir, err := os.Getwd()
		if err != nil {
			defaultDir = "."
		}
		Params.SetDefault("STATPING_DIR", defaultDir)
		Directory = defaultDir
	}
	Directory = Params.GetString("STATPING_DIR")
	Params.SetDefault("STATPING_DIR", Directory)
	Params.SetDefault("GO_ENV", "")
	Params.SetDefault("DISABLE_LOGS", false)
	Params.SetDefault("USE_ASSETS", false)
	Params.SetDefault("BASE_PATH", "")
	Params.SetDefault("ADMIN_USER", "admin")
	Params.SetDefault("ADMIN_PASSWORD", "admin")
	Params.SetDefault("MAX_OPEN_CONN", 25)
	Params.SetDefault("MAX_IDLE_CONN", 25)
	Params.SetDefault("MAX_LIFE_CONN", 5*time.Minute)
	Params.SetDefault("SAMPLE_DATA", true)
	Params.SetDefault("USE_CDN", false)
	Params.SetDefault("ALLOW_REPORTS", false)
	Params.SetDefault("POSTGRES_SSLMODE", "disable")
	Params.SetDefault("REMOVE_AFTER", 2160*time.Hour)
	Params.SetDefault("CLEANUP_INTERVAL", 1*time.Hour)

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
