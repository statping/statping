package configs

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
)

func LoadConfigFile(directory string) (*DbConfig, error) {
	var configs *DbConfig
	log.Infof("Attempting to read config file at: %s/config.yml ", directory)
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
