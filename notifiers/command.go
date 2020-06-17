package notifiers

import (
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

var _ notifier.Notifier = (*commandLine)(nil)

type commandLine struct {
	*notifications.Notification
}

func (c *commandLine) Select() *notifications.Notification {
	return c.Notification
}

var Command = &commandLine{&notifications.Notification{
	Method:      "command",
	Title:       "Command",
	Description: "Shell Command allows you to run a customized shell/bash Command on the local machine it's running on.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(1 * time.Second),
	Icon:        "fas fa-terminal",
	Host:        "/bin/bash",
	SuccessData: "curl -L http://localhost:8080",
	FailureData: "curl -L http://localhost:8080",
	DataType:    "text",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Executable Path",
		Placeholder: "/usr/bin/curl",
		DbField:     "host",
		SmallText:   "You can use '/bin/sh', '/bin/bash', '/usr/bin/curl' or an absolute path for an application.",
	}}},
}

func runCommand(app string, cmd ...string) (string, string, error) {
	utils.Log.Infof("Command notifier sending: %s %s", app, strings.Join(cmd, " "))
	outStr, errStr, err := utils.Command(app, cmd...)
	return outStr, errStr, err
}

// OnSuccess for commandLine will trigger successful service
func (c *commandLine) OnSuccess(s *services.Service) (string, error) {
	tmpl := ReplaceVars(c.SuccessData, s, nil)
	out, _, err := runCommand(c.Host, tmpl)
	return out, err
}

// OnFailure for commandLine will trigger failing service
func (c *commandLine) OnFailure(s *services.Service, f *failures.Failure) (string, error) {
	tmpl := ReplaceVars(c.FailureData, s, f)
	_, ouerr, err := runCommand(c.Host, tmpl)
	return ouerr, err
}

// OnTest for commandLine triggers when this notifier has been saved
func (c *commandLine) OnTest() (string, error) {
	tmpl := ReplaceVars(c.Var1, services.Example(true), failures.Example())
	in, out, err := runCommand(c.Host, tmpl)
	utils.Log.Infoln(in)
	utils.Log.Infoln(out)
	return out, err
}
