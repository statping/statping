package notifiers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/notifications"
	"github.com/statping-ng/statping-ng/types/notifier"
	"github.com/statping-ng/statping-ng/types/null"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/statping-ng/statping-ng/utils"
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
	Delay:       time.Duration(3 * time.Second),
	SuccessData: null.NewNullString(`{"id": "{{.Service.Id}}", "online": true}`),
	FailureData: null.NewNullString(`{"id": "{{.Service.Id}}", "online": false}`),
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
		Type:        "list",
		Title:       "HTTP Method",
		Placeholder: "POST",
		SmallText:   "Choose a HTTP method for example: GET, POST, DELETE, or PATCH.",
		DbField:     "Var1",
		Required:    true,
		ListOptions: []string{"GET", "POST", "PATCH", "DELETE"},
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

func (w *webhooker) Valid(values notifications.Values) error {
	return nil
}

func (w *webhooker) sendHttpWebhook(body string) (*http.Response, error) {
	client := new(http.Client)
	client.Timeout = 10 * time.Second
	req, err := http.NewRequest(w.Var1.String, w.Host.String, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	if w.ApiKey.String != "" {
		req.Header.Add("Content-Type", w.ApiKey.String)
	} else {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", "Statping-ng")
	req.Header.Set("Statping-Version", utils.Params.GetString("VERSION"))

	var customHeaders []string

	if w.ApiSecret.String != "" {
		customHeaders = strings.Split(w.ApiSecret.String, ",")
	} else {
		customHeaders = nil
	}

	for _, h := range customHeaders {
		keyVal := strings.SplitN(h, "=", 2)
		if len(keyVal) == 2 {
			if keyVal[0] != "" && keyVal[1] != "" {
				if strings.ToLower(keyVal[0]) == "host" {
					req.Host = strings.TrimSpace(keyVal[1])
				} else {
					req.Header.Set(keyVal[0], keyVal[1])
				}
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (w *webhooker) OnTest() (string, error) {
	f := failures.Example()
	s := services.Example(false)
	body := ReplaceVars(w.SuccessData.String, s, f)
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
func (w *webhooker) OnFailure(s services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(w.FailureData.String, s, f)
	resp, err := w.sendHttpWebhook(msg)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

// OnSuccess will trigger successful service
func (w *webhooker) OnSuccess(s services.Service) (string, error) {
	msg := ReplaceVars(w.SuccessData.String, s, failures.Failure{})
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
