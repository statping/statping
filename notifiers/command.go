// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	Title:       "Shell Command",
	Description: "Shell Command allows you to run a customized shell/bash Command on the local machine it's running on.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(1 * time.Second),
	Icon:        "fas fa-terminal",
	Host:        "/bin/bash",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Shell or Bash",
		Placeholder: "/bin/bash",
		DbField:     "host",
		SmallText:   "You can use '/bin/sh', '/bin/bash' or even an absolute path for an application.",
	}, {
		Type:        "text",
		Title:       "Command to Run on OnSuccess",
		Placeholder: "curl google.com",
		DbField:     "var1",
		SmallText:   "This Command will run every time a service is receiving a Successful event.",
	}, {
		Type:        "text",
		Title:       "Command to Run on OnFailure",
		Placeholder: "curl offline.com",
		DbField:     "var2",
		SmallText:   "This Command will run every time a service is receiving a Failing event.",
	}}},
}

func runCommand(app string, cmd ...string) (string, string, error) {
	outStr, errStr, err := utils.Command(app, cmd...)
	return outStr, errStr, err
}

// OnFailure for commandLine will trigger failing service
func (u *commandLine) OnFailure(s *services.Service, f *failures.Failure) error {
	msg := u.GetValue("var2")
	tmpl := ReplaceVars(msg, s, f)
	_, _, err := runCommand(u.Host, tmpl)
	return err
}

// OnSuccess for commandLine will trigger successful service
func (u *commandLine) OnSuccess(s *services.Service) error {
	msg := u.GetValue("var1")
	tmpl := ReplaceVars(msg, s, nil)
	_, _, err := runCommand(u.Host, tmpl)
	return err
}

// OnTest for commandLine triggers when this notifier has been saved
func (u *commandLine) OnTest() error {
	cmds := strings.Split(u.Var1, " ")
	in, out, err := runCommand(u.Host, cmds...)
	utils.Log.Infoln(in)
	utils.Log.Infoln(out)
	return err
}
