package configs

import (
	"errors"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadConfigs(cfgFile string) (*DbConfig, error) {
	writeAble, err := utils.DirWritable(utils.Directory)
	if err != nil {
		return nil, err
	}
	if !writeAble {
		return nil, errors.New("Directory %s is not writable: " + utils.Directory)
	}
	p := utils.Params
	log.Infof("Attempting to read config file at: %s", cfgFile)
	p.SetConfigFile(cfgFile)
	p.SetConfigType("yaml")
	p.ReadInConfig()

	db := new(DbConfig)
	content, err := utils.OpenFile(cfgFile)
	if err == nil {
		if err := yaml.Unmarshal([]byte(content), &db); err != nil {
			return nil, err
		}
	}

	if os.Getenv("DB_CONN") == "sqlite" || os.Getenv("DB_CONN") == "sqlite3" {
		db.DbConn = "sqlite3"
	}
	if db.DbConn != "" {
		p.Set("DB_CONN", db.DbConn)
	}
	if db.DbHost != "" {
		p.Set("DB_HOST", db.DbHost)
	}
	if db.DbPort != 0 {
		p.Set("DB_PORT", db.DbPort)
	}
	if db.DbPass != "" {
		p.Set("DB_PASS", db.DbPass)
	}
	if db.DbUser != "" {
		p.Set("DB_USER", db.DbUser)
	}
	if db.DbData != "" {
		p.Set("DB_DATABASE", db.DbData)
	}
	if db.Location != "" {
		p.Set("LOCATION", db.Location)
	}
	if db.ApiSecret != "" {
		p.Set("API_SECRET", db.ApiSecret)
	}
	if db.Language != "" {
		p.Set("LANGUAGE", db.Language)
	}
	if db.SendReports {
		p.Set("ALLOW_REPORTS", true)
	}
	if db.LetsEncryptEmail != "" {
		p.Set("LETSENCRYPT_EMAIL", db.LetsEncryptEmail)
	}
	if db.LetsEncryptHost != "" {
		p.Set("LETSENCRYPT_HOST", db.LetsEncryptHost)
	}
	if db.LetsEncryptEnable {
		p.Set("LETSENCRYPT_ENABLE", db.LetsEncryptEnable)
	}

	configs := &DbConfig{
		DbConn:            p.GetString("DB_CONN"),
		DbHost:            p.GetString("DB_HOST"),
		DbUser:            p.GetString("DB_USER"),
		DbPass:            p.GetString("DB_PASS"),
		DbData:            p.GetString("DB_DATABASE"),
		DbPort:            p.GetInt("DB_PORT"),
		Project:           p.GetString("NAME"),
		Description:       p.GetString("DESCRIPTION"),
		Domain:            p.GetString("DOMAIN"),
		Email:             p.GetString("EMAIL"),
		Username:          p.GetString("ADMIN_USER"),
		Password:          p.GetString("ADMIN_PASSWORD"),
		Location:          utils.Directory,
		SqlFile:           p.GetString("SQL_FILE"),
		Language:          p.GetString("LANGUAGE"),
		SendReports:       p.GetBool("ALLOW_REPORTS"),
		LetsEncryptEnable: p.GetBool("LETSENCRYPT_ENABLE"),
	}
	if configs.LetsEncryptEnable {
		configs.LetsEncryptHost = p.GetString("LETSENCRYPT_HOST")
		configs.LetsEncryptEmail = p.GetString("LETSENCRYPT_EMAIL")
	}
	log.WithFields(utils.ToFields(configs)).Debugln("read config file: " + cfgFile)

	if configs.DbConn == "" {
		return configs, errors.New("Starting in setup mode")
	}
	return configs, nil
}
