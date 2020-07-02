package notifiers

import (
	"github.com/statping/statping/types/errors"
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
	SuccessData: "/usr/bin/curl -L http://localhost:8080",
	FailureData: "/usr/bin/curl -L http://localhost:8080",
	DataType:    "text",
	Limits:      60,
}}

func runCommand(cmd string) (string, string, error) {
	utils.Log.Infof("Command notifier sending: %s", cmd)
	cmdApp := strings.Split(cmd, " ")
	if len(cmd) == 0 {
		return "", "", errors.New("you need at least 1 command")
	}
	var cmdArgs []string
	if len(cmd) > 1 {
		cmdArgs = append(cmdArgs, cmd[1:])
	}
	outStr, errStr, err := utils.Command(cmdApp[0], cmdArgs...)
	return outStr, errStr, err
}

// OnSuccess for commandLine will trigger successful service
func (c *commandLine) OnSuccess(s services.Service) (string, error) {
	tmpl := ReplaceVars(c.SuccessData, s, failures.Failure{})
	out, _, err := runCommand(tmpl)
	return out, err
}

// OnFailure for commandLine will trigger failing service
func (c *commandLine) OnFailure(s services.Service, f failures.Failure) (string, error) {
	tmpl := ReplaceVars(c.FailureData, s, f)
	out, _, err := runCommand(tmpl)
	return out, err
}

// OnTest for commandLine triggers when this notifier has been saved
func (c *commandLine) OnTest() (string, error) {
	tmpl := ReplaceVars(c.Var1, services.Example(true), failures.Example())
	in, out, err := runCommand(tmpl)
	utils.Log.Infoln(in)
	utils.Log.Infoln(out)
	return out, err
}

// OnSave will trigger when this notifier is saved
func (c *commandLine) OnSave() (string, error) {
	return "", nil
}
