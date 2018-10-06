// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

const (
	SLACK_METHOD     = "slack"
	FAILING_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is currently failing", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification with a HTTP Status code of {{.Service.LastStatusCode}}.", "fields": [ { "title": "Expected", "value": "{{.Service.Expected}}", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#FF0000", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	SUCCESS_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is now back online", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification.", "fields": [ { "title": "Issue", "value": "Awesome Project", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#00FF00", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	SLACK_TEXT       = `{"text":"{{.}}"}`
)

type slack struct {
	*notifier.Notification
}

var slacker = &slack{&notifier.Notification{
	Method:      SLACK_METHOD,
	Title:       "slack",
	Description: "Send notifications to your slack channel when a service is offline. Insert your Incoming webhooker URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Host:        "https://webhooksurl.slack.com/***",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Incoming webhooker Url",
		Placeholder: "Insert your slack webhook URL here.",
		SmallText:   "Incoming webhooker URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">slack Apps</a>",
		DbField:     "Host",
		Required:    true,
	}}},
}

func parseSlackMessage(temp string, data interface{}) error {
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	err := slackTemp.Execute(buf, data)
	if err != nil {
		return err
	}
	slacker.AddQueue(buf.String())
	return nil
}

type slackMessage struct {
	Service  *types.Service
	Template string
	Time     int64
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(slacker)
	if err != nil {
		panic(err)
	}
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (u *slack) Send(msg interface{}) error {
	message := msg.(string)
	client := new(http.Client)
	res, err := client.Post(u.Host, "application/json", bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	//contents, _ := ioutil.ReadAll(res.Body)
	return nil
}

func (u *slack) Select() *notifier.Notification {
	return u.Notification
}

func (u *slack) OnTest() error {
	client := new(http.Client)
	res, err := client.Post(u.Host, "application/json", bytes.NewBuffer([]byte(`{"text":"testing message"}`)))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)
	if string(contents) != "ok" {
		return errors.New("The slack response was incorrect, check the URL")
	}
	return err
}

// OnFailure will trigger failing service
func (u *slack) OnFailure(s *types.Service, f *types.Failure) {
	message := slackMessage{
		Service:  s,
		Template: FAILING_TEMPLATE,
		Time:     time.Now().Unix(),
	}
	parseSlackMessage(FAILING_TEMPLATE, message)
	u.Online = false
}

// OnSuccess will trigger successful service
func (u *slack) OnSuccess(s *types.Service) {
	if !u.Online {
		message := slackMessage{
			Service:  s,
			Template: SUCCESS_TEMPLATE,
			Time:     time.Now().Unix(),
		}
		parseSlackMessage(SUCCESS_TEMPLATE, message)
	}
	u.Online = true
}

// OnSave triggers when this notifier has been saved
func (u *slack) OnSave() error {
	message := fmt.Sprintf("Notification %v is receiving updated information.", u.Method)
	u.AddQueue(message)
	return nil
}
