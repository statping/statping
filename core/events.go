package core

import (
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/hunterlong/statup/notifications"
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
	if notifications.SlackComm != nil {
		onFailureSlack(s, f)
	}
	if notifications.EmailComm != nil {
		onFailureEmail(s, f)
	}
	if notifications.PushoverComm != nil {
		onFailurePushover(s, f)
	}
}

func onFailureSlack(s *Service, f FailureData) {
	slack := SelectCommunication(2)
	if slack.Enabled {
		msg := fmt.Sprintf("Service %v is currently offline! Issue: %v", s.Name, f.Issue)
		notifications.SendSlack(msg)
	}
}

func onFailurePushover(s *Service, f FailureData) {
	pushover := SelectCommunication(3)
	if pushover.Enabled {
		// Trick to send only one notification when service is down
		if time.Now().Sub(s.LastOnline).Seconds() < float64(2*s.Interval) {
			admin, _ := SelectUser(1)
			if admin.PushoverUserKey != "" {
				notification := &types.PushoverNotification{
					To:      admin.PushoverUserKey,
					Message: fmt.Sprintf("Service %v is currently offline! Issue: %v", s.Name, f.Issue),
				}
				notifications.SendPushover(notification)
			}
		}
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
		notifications.SendEmail(EmailBox, email)
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
