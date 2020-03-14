package main

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/services"
)

func InitApp() error {
	if _, err := core.Select(); err != nil {
		return err
	}

	if _, err := services.SelectAllServices(true); err != nil {
		return err
	}

	go services.CheckServices()

	notifiers.InitNotifiers()

	database.StartMaintenceRoutine()
	core.App.Setup = true
	return nil
}
