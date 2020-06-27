package notifiers

import (
	"bytes"
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var _ notifier.Notifier = (*webhooker)(nil)

const (
	webhookMethod = "webhook"
)

type webhooker struct {
	*notifications.Notification
}

var Webhook = &webhooker{&notifications.Notification{
	Method:      webhookMethod,
	Title:       "Webhook",
	Description: "Send a custom HTTP request to a specific URL with your own body, headers, and parameters.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "fas fa-code-branch",
	Delay:       time.Duration(1 * time.Second),
	SuccessData: `{"id": {{.Service.Id}}, "online": true}`,
	FailureData: `{"id": {{.Service.Id}}, "online": false}`,
	DataType:    "json",
	Limits:      180,
	Form: []notifications.NotificationForm{{
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

func (w *webhooker) Select() *notifications.Notification {
	return w.Notification
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

func (w *webhooker) OnTest() (string, error) {
	f := failures.Example()
	s := services.Example(false)
	body := ReplaceVars(w.SuccessData, s, f)
	resp, err := w.sendHttpWebhook(body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	out := fmt.Sprintf("Webhook notifier received: '%v'", string(content))
	utils.Log.Infoln(out)
	return out, err
}

// OnFailure will trigger failing service
func (w *webhooker) OnFailure(s *services.Service, f *failures.Failure) (string, error) {
	msg := ReplaceVars(w.FailureData, s, f)
	resp, err := w.sendHttpWebhook(msg)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

// OnSuccess will trigger successful service
func (w *webhooker) OnSuccess(s *services.Service) (string, error) {
	msg := ReplaceVars(w.SuccessData, s, nil)
	resp, err := w.sendHttpWebhook(msg)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

// OnSave will trigger when this notifier is saved
func (w *webhooker) OnSave() (string, error) {
	return "", nil
}
