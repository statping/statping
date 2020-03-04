package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/notifiers"
	"github.com/hunterlong/statping/types/services"
)

func InitApp() error {
	if _, err := Select(); err != nil {
		return err
	}
	//if err := InsertNotifierDB(); err != nil {
	//	return err
	//}
	//if err := InsertIntegratorDB(); err != nil {
	//	return err
	//}
	if _, err := services.SelectAllServices(true); err != nil {
		return err
	}
	if err := notifiers.AttachNotifiers(); err != nil {
		return err
	}
	//App.Notifications = notifications.AllCommunications
	//if err := integrations.AddIntegrations(); err != nil {
	//	return err
	//}
	//App.Integrations = integrations.Integrations

	go services.CheckServices()

	database.StartMaintenceRoutine()
	App.Setup = true
	return nil
}
