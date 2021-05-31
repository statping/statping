package configs

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"os"
	"path/filepath"
	"strings"
)

var log = utils.Log.WithField("type", "configs")

// ConnectConfigs will connect to the database and save the config.yml file
func ConnectConfigs(configs *DbConfig, retry bool) error {
	err := Connect(configs, retry)
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}
	if err := configs.Save(utils.Directory); err != nil {
		return errors.Wrap(err, "error saving configuration")
	}
	return nil
}

// findDbFile will attempt to find the "statping.db" database file in the current
// working directory, or from STATPING_DIR env.
func findDbFile(configs *DbConfig) (string, error) {
	location := utils.Directory + "/" + SqliteFilename
	if configs == nil {
		file, err := findSQLin(utils.Directory)
		if err != nil {
			log.Errorln(err)
			return location, nil
		}
		location = file
	}
	if configs != nil && configs.SqlFile != "" {
		return configs.SqlFile, nil
	}
	return location, nil
}

// findSQLin walks the current walking directory for statping.db
func findSQLin(path string) (string, error) {
	filename := SqliteFilename
	var found []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".db" {
			filename = info.Name()
			found = append(found, filename)
		}
		return nil
	})
	if err != nil {
		return filename, err
	}
	if len(found) > 1 {
		return filename, errors.Errorf("found multiple database files: %s", strings.Join(found, ", "))
	}
	return filename, nil
}
