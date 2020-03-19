package configs

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
)

func loadConfigEnvs() (*DbConfig, error) {
	var err error

	log.Infof("Loading configs from environment variables")

	loadDotEnvs()

	dbConn := utils.Getenv("DB_CONN", "").(string)
	dbHost := utils.Getenv("DB_HOST", "").(string)
	dbUser := utils.Getenv("DB_USER", "").(string)
	dbPass := utils.Getenv("DB_PASS", "").(string)
	dbData := utils.Getenv("DB_DATABASE", "").(string)
	dbPort := utils.Getenv("DB_PORT", defaultPort(dbConn)).(int)
	name := utils.Getenv("NAME", "Statping").(string)
	desc := utils.Getenv("DESCRIPTION", "Statping Monitoring Sample Data").(string)
	user := utils.Getenv("ADMIN_USER", "admin").(string)
	password := utils.Getenv("ADMIN_PASS", "admin").(string)
	domain := utils.Getenv("DOMAIN", "").(string)
	sqlFile := utils.Getenv("SQL_FILE", "").(string)

	if dbConn != "" && dbConn != "sqlite" {
		if dbHost == "" {
			return nil, errors.New("Missing DB_HOST environment variable")
		}
		if dbPort == 0 {
			return nil, errors.New("Missing DB_PORT environment variable")
		}
		if dbUser == "" {
			return nil, errors.New("Missing DB_USER environment variable")
		}
		if dbPass == "" {
			return nil, errors.New("Missing DB_PASS environment variable")
		}
		if dbData == "" {
			return nil, errors.New("Missing DB_DATABASE environment variable")
		}
	}

	config := &DbConfig{
		DbConn:      dbConn,
		DbHost:      dbHost,
		DbUser:      dbUser,
		DbPass:      dbPass,
		DbData:      dbData,
		DbPort:      dbPort,
		Project:     name,
		Description: desc,
		Domain:      domain,
		Email:       "",
		Username:    user,
		Password:    password,
		Error:       nil,
		Location:    utils.Directory,
		SqlFile:     sqlFile,
	}
	return config, err
}

// loadDotEnvs attempts to load database configs from a '.env' file in root directory
func loadDotEnvs() {
	err := godotenv.Overload(utils.Directory + "/" + ".env")
	if err == nil {
		log.Warnln("Environment file '.env' found")
		envs, _ := godotenv.Read(utils.Directory + "/" + ".env")
		for k, e := range envs {
			log.Infof("Overwriting %s=%s\n", k, e)
		}
		log.Warnln("These environment variables will overwrite any existing")
	}
}
