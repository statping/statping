package core

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
	"upper.io/db.v3/lib/sqlbuilder"
)

func OnLoad(db sqlbuilder.Database) {
	for _, p := range CoreApp.AllPlugins {
		p.OnLoad(db)
	}
}

func OnSuccess(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSuccess(structs.Map(s))
	}
}

func OnFailure(s *Service, f FailureData) {
	for _, p := range CoreApp.AllPlugins {
		p.OnFailure(structs.Map(s))
	}
	if notifiers.SendSlack("im failing") != nil {
		onFailureSlack(s, f)
	}
	if notifiers.EmailComm != nil {
		onFailureEmail(s, f)
	}
}

func onFailureSlack(s *Service, f FailureData) {
	slack := SelectCommunication(2)
	if slack.Enabled {
		//msg := fmt.Sprintf("Service %v is currently offline! Issue: %v", s.Name, f.Issue)
		//communications.SendSlack(msg)
	}
}

type failedEmail struct {
	Service     *Service
	FailureData FailureData
	Domain      string
}

func onFailureEmail(s *Service, f FailureData) {
	email := SelectCommunication(1)
	if email.Enabled {
		data := failedEmail{s, f, CoreApp.Domain}
		admin, _ := SelectUser(1)
		email := &types.Email{
			To:       admin.Email,
			Subject:  fmt.Sprintf("Service %v is Down", s.Name),
			Template: "failure.html",
			Data:     data,
			From:     email.Var1,
		}
		notifiers.SendEmail(EmailBox, email)
	}
}

func OnSettingsSaved(c *Core) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSettingsSaved(structs.Map(c))
	}
}

func OnNewUser(u *User) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewUser(structs.Map(u))
	}
}

func OnNewService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewService(structs.Map(s))
	}
}

func OnDeletedService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnDeletedService(structs.Map(s))
	}
}

func OnUpdateService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnUpdatedService(structs.Map(s))
	}
}

func SelectPlugin(name string) plugin.PluginActions {
	for _, p := range CoreApp.AllPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return plugin.PluginInfo{}
}
