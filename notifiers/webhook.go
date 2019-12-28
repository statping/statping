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
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	webhookMethod = "Webhook"
)

type webhooker struct {
	*notifier.Notification
}

var Webhook = &webhooker{&notifier.Notification{
	Method:      webhookMethod,
	Title:       "HTTP webhooker",
	Description: "Send a custom HTTP request to a specific URL with your own body, headers, and parameters.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "fas fa-code-branch",
	Delay:       time.Duration(1 * time.Second),
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "HTTP Endpoint",
		Placeholder: "http://webhookurl.com/JW2MCP4SKQP",
		SmallText:   "Insert the URL for your HTTP Requests.",
		DbField:     "Host",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "HTTP Method",
		Placeholder: "POST",
		SmallText:   "Choose a HTTP method for example: GET, POST, DELETE, or PATCH.",
		DbField:     "Var1",
		Required:    true,
	}, {
		Type:        "textarea",
		Title:       "HTTP Body",
		Placeholder: `{"service_id": "%s.Id", "service_name": "%s.Name"}`,
		SmallText:   "Optional HTTP body for a POST request. You can insert variables into your body request.<br>%service.Id, %service.Name, %service.Online<br>%failure.Issue",
		DbField:     "Var2",
	}, {
		Type:        "text",
		Title:       "Content Type",
		Placeholder: `application/json`,
		SmallText:   "Optional content type for example: application/json or text/plain",
		DbField:     "api_key",
	}, {
		Type:        "text",
		Title:       "Header",
		Placeholder: "Authorization=Token12345",
		SmallText:   "Optional Headers for request use format: KEY=Value,Key=Value",
		DbField:     "api_secret",
	},
	}}}

// Send will send a HTTP Post to the webhooker API. It accepts type: string
func (w *webhooker) Send(msg interface{}) error {
	resp, err := w.sendHttpWebhook(msg.(string))
	if err == nil {
		resp.Body.Close()
	}
	return err
}

func (w *webhooker) Select() *notifier.Notification {
	return w.Notification
}

func replaceBodyText(body string, s *types.Service, f *types.Failure) string {
	body = utils.ConvertInterface(body, s)
	body = utils.ConvertInterface(body, f)
	return body
}

func (w *webhooker) sendHttpWebhook(body string) (*http.Response, error) {
	utils.Log.Infoln(fmt.Sprintf("sending body: '%v' to %v as a %v request", body, w.Host, w.Var1))
	client := new(http.Client)
	client.Timeout = time.Duration(10 * time.Second)
	var buf *bytes.Buffer
	buf = bytes.NewBuffer(nil)
	if w.Var2 != "" {
		buf = bytes.NewBuffer([]byte(body))
	}
	req, err := http.NewRequest(w.Var1, w.Host, buf)
	if err != nil {
		return nil, err
	}
	if w.ApiSecret != "" {
		splitArray := strings.Split(w.ApiSecret, ",")
		for _, a := range splitArray {
			split := strings.Split(a, "=")
			req.Header.Add(split[0], split[1])
		}
	}
	if w.ApiKey != "" {
		req.Header.Add("Content-Type", w.ApiKey)
	}
	req.Header.Set("User-Agent", "Statping")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (w *webhooker) OnTest() error {
	body := replaceBodyText(w.Var2, notifier.ExampleService, nil)
	resp, err := w.sendHttpWebhook(body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	utils.Log.Infoln(fmt.Sprintf("Webhook notifier received: '%v'", string(content)))
	return err
}

// OnFailure will trigger failing service
func (w *webhooker) OnFailure(s *types.Service, f *types.Failure) {
	msg := replaceBodyText(w.Var2, s, f)
	w.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnSuccess will trigger successful service
func (w *webhooker) OnSuccess(s *types.Service) {
	if !s.Online {
		w.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		msg := replaceBodyText(w.Var2, s, nil)
		w.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
	}
}

// OnSave triggers when this notifier has been saved
func (w *webhooker) OnSave() error {
	return nil
}
