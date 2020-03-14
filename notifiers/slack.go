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
	"bytes"
	"errors"
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"text/template"
	"time"
)

var _ notifier.Notifier = (*slack)(nil)

const (
	slackMethod     = "slack"
	failingTemplate = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is currently failing", "text": "Your Statping service <{{.Service.Domain}}|{{.Service.Name}}> has just received a Failure notification based on your expected results. {{.Service.Name}} responded with a HTTP Status code of {{.Service.LastStatusCode}}.", "fields": [ { "title": "Expected Status Code", "value": "{{.Service.ExpectedStatus}}", "short": true }, { "title": "Received Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ,{ "title": "Error Message", "value": "{{.Issue}}", "short": false } ], "color": "#FF0000", "thumb_url": "https://statping.com", "footer": "Statping", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	successTemplate = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is now back online", "text": "Your Statping service <{{.Service.Domain}}|{{.Service.Name}}> is now back online and meets your expected responses.", "color": "#00FF00", "thumb_url": "https://statping.com", "footer": "Statping", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	slackText       = `{"text":"{{.}}"}`
)

type slack struct {
	*notifications.Notification
}

func (s *slack) Select() *notifications.Notification {
	return s.Notification
}

var slacker = &slack{&notifications.Notification{
	Method:      slackMethod,
	Title:       "slack",
	Description: "Send notifications to your slack channel when a service is offline. Insert your Incoming webhooker URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Host:        "https://webhooksurl.slack.com/***",
	Icon:        "fab fa-slack",
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Incoming webhooker Url",
		Placeholder: "Insert your slack Webhook URL here.",
		SmallText:   "Incoming webhooker URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">slack Apps</a>",
		DbField:     "Host",
		Required:    true,
	}}},
}

func parseSlackMessage(id int64, temp string, data interface{}) string {
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	err := slackTemp.Execute(buf, data)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

type slackMessage struct {
	Service  *services.Service
	Template string
	Time     int64
	Issue    string
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (u *slack) sendSlack(msg string) error {
	contents, resp, err := utils.HttpRequest(u.Host, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true)
	defer resp.Body.Close()
	fmt.Println("CONTENTS: ", string(contents))
	return err
}

func (u *slack) OnTest() error {
	contents, resp, err := utils.HttpRequest(u.Host, "POST", "application/json", nil, bytes.NewBuffer([]byte(`{"text":"testing message"}`)), time.Duration(10*time.Second), true)
	defer resp.Body.Close()
	if string(contents) != "ok" {
		return errors.New("the slack response was incorrect, check the URL")
	}
	return err
}

// OnFailure will trigger failing service
func (u *slack) OnFailure(s *services.Service, f *failures.Failure) error {
	message := slackMessage{
		Service:  s,
		Template: failingTemplate,
		Time:     utils.Now().Unix(),
	}
	msg := parseSlackMessage(s.Id, failingTemplate, message)
	return u.sendSlack(msg)
}

// OnSuccess will trigger successful service
func (u *slack) OnSuccess(s *services.Service) error {
	message := slackMessage{
		Service:  s,
		Template: successTemplate,
		Time:     utils.Now().Unix(),
	}
	msg := parseSlackMessage(s.Id, successTemplate, message)

	fmt.Println("Sending OnSuccess message!")
	fmt.Println(msg)
	fmt.Printf("%s\n", u.Host)
	return u.sendSlack(msg)
}
