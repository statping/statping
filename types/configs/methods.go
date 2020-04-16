package configs

import "github.com/statping/statping/utils"

// Save will initially create the config.yml file
func (d *DbConfig) Save(directory string) error {
	if err := utils.Params.SafeWriteConfigAs(directory + "/config.yml"); err != nil {
		return nil
	}
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
