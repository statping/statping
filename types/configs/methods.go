package configs

import "github.com/statping/statping/utils"

// Save will initially create the config.yml file
func (d *DbConfig) Save(directory string) error {
	if err := utils.Params.WriteConfigAs(directory + "/config.yml"); err != nil {
		return nil
	}
	return nil
}
