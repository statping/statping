package configs

import (
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

var log = utils.Log

func ConnectConfigs(configs *DbConfig) error {
	err := Connect(configs, true)
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}
	if err := configs.Save(utils.Directory); err != nil {
		return errors.Wrap(err, "error saving configuration")
	}

	exists := database.DB().HasTable("core")
	if !exists {
		return InitialSetup(configs)
	}
	return nil
}

func LoadConfigs() (*DbConfig, error) {
	writeAble, err := utils.DirWritable(utils.Directory)
	if err != nil {
		return nil, err
	}
	if !writeAble {
		return nil, errors.Errorf("Directory %s is not writable!", utils.Directory)
	}

	dbConn := utils.Getenv("DB_CONN", "").(string)
	if dbConn != "" {
		configs, err := loadConfigEnvs()
		if err != nil {
			return LoadConfigFile(utils.Directory)
		}
		return configs, nil
	}

	return LoadConfigFile(utils.Directory)
}

func findDbFile(configs *DbConfig) string {
	if configs == nil {
		return findSQLin(utils.Directory)
	}
	if configs.SqlFile != "" {
		return configs.SqlFile
	}
	return utils.Directory + "/" + SqliteFilename
}

func findSQLin(path string) string {
	filename := SqliteFilename
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".db" {
			fmt.Println("DB file is now: ", info.Name())
			filename = info.Name()
		}
		return nil
	})
	if err != nil {
		log.Error(err)
	}
	return filename
}
