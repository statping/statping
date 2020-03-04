package configs

import (
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
)

func loadConfigFile(directory string) (*DbConfig, error) {
	var configs *DbConfig

	log.Debugln("Attempting to read config file at: " + directory + "/config.yml")
	file, err := utils.OpenFile(directory + "/config.yml")
	if err != nil {
		core.App.Setup = false
		return nil, errors.Wrapf(err, "config.yml file not found at %s/config.yml - starting in setup mode", directory)
	}
	err = yaml.Unmarshal([]byte(file), &configs)
	if err != nil {
		return nil, errors.Wrap(err, "yaml file not formatted correctly")
	}
	log.WithFields(utils.ToFields(configs)).Debugln("read config file: " + directory + "/config.yml")

	return configs, nil
}
