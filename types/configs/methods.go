package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Save will initially create the config.yml file
func (d *DbConfig) Save(directory string) error {
	data, err := yaml.Marshal(d)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(directory+"/config.yml", data, os.ModePerm); err != nil {
		return err
	}
	d.filename = directory + "/config.yml"
	return nil
}

// defaultPort accepts a database type and returns its default port
func defaultPort(db string) int {
	switch db {
	case "mysql":
		return 3306
	case "postgres":
		return 5432
	case "mssql":
		return 1433
	default:
		return 0
	}
}
