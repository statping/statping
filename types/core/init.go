package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/services"
)

func InitApp() error {
	if _, err := Select(); err != nil {
		return err
	}

	if _, err := services.SelectAllServices(true); err != nil {
		return err
	}

	go services.CheckServices()

	database.StartMaintenceRoutine()
	App.Setup = true
	return nil
}
