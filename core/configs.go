package core

import (
	"errors"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/types"
	"io/ioutil"
	"os"
)

func LoadConfig() (*types.Config, error) {
	if os.Getenv("DB_CONN") != "" {
		return LoadUsingEnv()
	}
	Configs = new(types.Config)
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, errors.New("config.yml file not found - starting in setup mode")
	}
	err = yaml.Unmarshal(file, &Configs)
	CoreApp.DbConnection = Configs.Connection
	return Configs, err
}

func LoadUsingEnv() (*types.Config, error) {
	Configs = new(types.Config)
	if os.Getenv("DB_CONN") == "" {
		return nil, errors.New("Missing DB_CONN environment variable")
	}
	if os.Getenv("DB_HOST") == "" {
		return nil, errors.New("Missing DB_HOST environment variable")
	}
	if os.Getenv("DB_USER") == "" {
		return nil, errors.New("Missing DB_USER environment variable")
	}
	if os.Getenv("DB_PASS") == "" {
		return nil, errors.New("Missing DB_PASS environment variable")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return nil, errors.New("Missing DB_DATABASE environment variable")
	}
	Configs.Connection = os.Getenv("DB_CONN")
	Configs.Host = os.Getenv("DB_HOST")
	Configs.Port = os.Getenv("DB_PORT")
	Configs.User = os.Getenv("DB_USER")
	Configs.Password = os.Getenv("DB_PASS")
	Configs.Database = os.Getenv("DB_DATABASE")
	CoreApp.DbConnection = os.Getenv("DB_CONN")
	CoreApp.Name = os.Getenv("NAME")
	CoreApp.Domain = os.Getenv("DOMAIN")
	if os.Getenv("USE_CDN") == "true" {
		CoreApp.UseCdn = true
	}
	return Configs, nil
}

func ifOr(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
