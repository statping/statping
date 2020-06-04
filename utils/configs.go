package utils

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

var (
	Params    *viper.Viper
	configLog = Log.WithField("type", "configs")
)

func initCLI() {
	Params = viper.New()
	Params.AutomaticEnv()
	setDefaults()
	Directory = Params.GetString("STATPING_DIR")
	//Params.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	Params.SetConfigName("config")
	Params.SetConfigType("yml")
	Params.AddConfigPath(Directory)

	Params.ReadInConfig()

	Params.AddConfigPath(Directory)
	Params.SetConfigFile(".env")
	Params.ReadInConfig()

	Params.Set("VERSION", version)

	// check if logs are disabled
	if !Params.GetBool("DISABLE_LOGS") {
		Log.Out = ioutil.Discard

		Log.Debugln("current working directory: ", Directory)
		Log.AddHook(new(hook))
		Log.SetNoLock()
		checkVerboseMode()
	}
}

func setDefaults() {
	var err error
	defaultDir, err := os.Getwd()
	if err != nil {
		configLog.Errorln(err)
		defaultDir = "."
	}
	Params.SetDefault("STATPING_DIR", defaultDir)
	Params.SetDefault("GO_ENV", "")
	Params.SetDefault("DB_CONN", "")
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
	Params.SetDefault("NAME", "Statping Sample Data")
	Params.SetDefault("DOMAIN", "http://localhost:8080")
	Params.SetDefault("DESCRIPTION", "This status page has sample data included")
	Params.SetDefault("REMOVE_AFTER", 2160*time.Hour)
	Params.SetDefault("CLEANUP_INTERVAL", 1*time.Hour)
	Params.SetDefault("LANGUAGE", "en")
	Params.SetDefault("LOGS_MAX_COUNT", 5)
	Params.SetDefault("LOGS_MAX_AGE", 28)
	Params.SetDefault("LOGS_MAX_SIZE", 16)

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
