// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"bytes"
	"errors"
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"strings"
	"text/template"
	"time"
)

const (
	slackMethod     = "slack"
	failingTemplate = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is currently failing", "text": "Your Statping service <{{.Service.Domain}}|{{.Service.Name}}> has just received a Failure notification based on your expected results. {{.Service.Name}} responded with a HTTP Status code of {{.Service.LastStatusCode}}.", "fields": [ { "title": "Expected Status Code", "value": "{{.Service.ExpectedStatus}}", "short": true }, { "title": "Received Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ,{ "title": "Error Message", "value": "{{.Issue}}", "short": false } ], "color": "#FF0000", "thumb_url": "https://statping.com", "footer": "Statping", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	successTemplate = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is now back online", "text": "Your Statping service <{{.Service.Domain}}|{{.Service.Name}}> is now back online and meets your expected responses.", "color": "#00FF00", "thumb_url": "https://statping.com", "footer": "Statping", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	slackText       = `{"text":"{{.}}"}`
)

type slack struct {
	*notifier.Notification
}

var Slacker = &slack{&notifier.Notification{
	Method:      slackMethod,
	Title:       "slack",
	Description: "Send notifications to your slack channel when a service is offline. Insert your Incoming webhooker URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Host:        "https://webhooksurl.slack.com/***",
	Icon:        "fab fa-slack",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Incoming webhooker Url",
		Placeholder: "Insert your slack Webhook URL here.",
		SmallText:   "Incoming webhooker URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">slack Apps</a>",
		DbField:     "Host",
		Required:    true,
	}}},
}

func parseSlackMessage(id int64, temp string, data interface{}) error {
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	err := slackTemp.Execute(buf, data)
	if err != nil {
		return err
	}
	Slacker.AddQueue(fmt.Sprintf("service_%v", id), buf.String())
	return nil
}

type slackMessage struct {
	Service  *types.Service
	Template string
	Time     int64
	Issue    string
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (u *slack) Send(msg interface{}) error {
	message := msg.(string)
	_, _, err := utils.HttpRequest(u.Host, "POST", "application/json", nil, strings.NewReader(message), time.Duration(10*time.Second), true)
	return err
}

func (u *slack) Select() *notifier.Notification {
	return u.Notification
}

func (u *slack) OnTest() error {
	contents, _, err := utils.HttpRequest(u.Host, "POST", "application/json", nil, bytes.NewBuffer([]byte(`{"text":"testing message"}`)), time.Duration(10*time.Second), true)
	if string(contents) != "ok" {
		return errors.New("The slack response was incorrect, check the URL")
	}
	return err
}

// OnFailure will trigger failing service
func (u *slack) OnFailure(s *types.Service, f *types.Failure) {
	message := slackMessage{
		Service:  s,
		Template: failingTemplate,
		Time:     utils.Now().Unix(),
		Issue:    f.Issue,
	}
	parseSlackMessage(s.Id, failingTemplate, message)
}

// OnSuccess will trigger successful service
func (u *slack) OnSuccess(s *types.Service) {
	if !s.Online {
		u.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		message := slackMessage{
			Service:  s,
			Template: successTemplate,
			Time:     utils.Now().Unix(),
		}
		parseSlackMessage(s.Id, successTemplate, message)
	}
}

// OnSave triggers when this notifier has been saved
func (u *slack) OnSave() error {
	message := fmt.Sprintf("Notification %v is receiving updated information.", u.Method)
	u.AddQueue("saved", message)
	return nil
}
