package utils

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

var (
	Params *viper.Viper
)

func InitEnvs() {
	if Params != nil {
		return
	}
	Params = viper.New()
	Params.AutomaticEnv()

	var err error
	defaultDir, err := os.Getwd()
	if err != nil {
		Log.Errorln(err)
		defaultDir = "."
	}
	Params.SetDefault("DISABLE_HTTP", false)
	Params.SetDefault("STATPING_DIR", defaultDir)
	Params.SetDefault("GO_ENV", "production")
	Params.SetDefault("DEBUG", false)
	Params.SetDefault("DEMO_MODE", false)
	Params.SetDefault("DB_CONN", "")
	Params.SetDefault("DB_DSN", "")
	Params.SetDefault("DISABLE_LOGS", false)
	Params.SetDefault("USE_ASSETS", false)
	Params.SetDefault("BASE_PATH", "")
	Params.SetDefault("ADMIN_USER", "admin")
	Params.SetDefault("ADMIN_PASSWORD", "admin")
	Params.SetDefault("ADMIN_EMAIL", "info@admin.com")
	Params.SetDefault("MAX_OPEN_CONN", 25)
	Params.SetDefault("MAX_IDLE_CONN", 25)
	Params.SetDefault("MAX_LIFE_CONN", 5*time.Minute)
	Params.SetDefault("SAMPLE_DATA", true)
	Params.SetDefault("USE_CDN", false)
	Params.SetDefault("ALLOW_REPORTS", true)
	Params.SetDefault("POSTGRES_SSLMODE", "disable")
	Params.SetDefault("NAME", "Statping Sample Data")
	Params.SetDefault("DOMAIN", "http://localhost:8080")
	Params.SetDefault("DESCRIPTION", "This status page has sample data included")
	Params.SetDefault("REMOVE_AFTER", 2160*time.Hour)
	Params.SetDefault("CLEANUP_INTERVAL", 1*time.Hour)
	Params.SetDefault("LANGUAGE", "en")
	Params.SetDefault("LETSENCRYPT_HOST", "")
	Params.SetDefault("LETSENCRYPT_EMAIL", "")
	Params.SetDefault("LETSENCRYPT_LOCAL", false)
	Params.SetDefault("LETSENCRYPT_ENABLE", false)
	Params.SetDefault("READ_ONLY", false)
	Params.SetDefault("LOGS_MAX_COUNT", 5)
	Params.SetDefault("LOGS_MAX_AGE", 28)
	Params.SetDefault("LOGS_MAX_SIZE", 16)
	Params.SetDefault("DISABLE_COLORS", false)

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

	Directory = Params.GetString("STATPING_DIR")
	//Params.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	Params.SetConfigName("config")
	Params.SetConfigType("yml")
	Params.AddConfigPath(Directory)
	Params.ReadInConfig()

	Params.AddConfigPath(Directory)
	Params.SetConfigFile(".env")
	Params.ReadInConfig()

	// check if logs are disabled
	if Params.GetBool("DISABLE_LOGS") {
		Log.Out = ioutil.Discard
		return
	}
	Log.Debugln("current working directory: ", Directory)
	Log.AddHook(new(hook))
	Log.SetNoLock()
	checkVerboseMode()
}
