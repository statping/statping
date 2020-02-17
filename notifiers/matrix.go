// Statping
// Copyright (C) 2020.  Hunter Long and the project contributors
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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"io"
	"net/http"
	"time"
)

type matrix struct {
	*notifier.Notification
}

var Matrix = &matrix{
	Notification: &notifier.Notification{
		Method:      "matrix",
		Title:       "Matrix",
		Description: "Receive notifications on your Matrix room when a service has an issue. You must get a username and token. It is recommended to create a dedicated user for this.",
		Author:      "Tero Vierimaa",
		AuthorUrl:   "https://github.com/tryffel",
		Icon:        "far fa-bell",
		Delay:       5 * time.Second,
		Form: []notifier.NotificationForm{{
			Type:        "text",
			Title:       "Homeserver url",
			Placeholder: "https://matrix.org",
			SmallText:   "Enter homeserver url, http(s) included",
			DbField:     "host",
			Required:    true,
		}, {
			Type:        "text",
			Title:       "Room Id",
			Placeholder: "!MzCSpOkYukLxYisAbY:matrix.org",
			SmallText:   "Enter full room id",
			DbField:     "var1",
			Required:    true,
		}, {
			Type:        "text",
			Title:       "Token",
			Placeholder: "MDAxyz...",
			SmallText:   "Enter user token",
			DbField:     "api_secret",
			Required:    true,
		}, {
			Type:        "text",
			Title:       "Event type",
			Placeholder: "m.text",
			SmallText:   "Either m.text (normal message) or m.notice (silent message)",
			DbField:     "var2",
			Required:    false,
		}},
	},
}

// matrixUrl, params: room, nonce
var matrixUrlPath = "/_matrix/client/r0/rooms/%s/send/m.room.message/%d"

func (m *matrix) Select() *notifier.Notification {
	return m.Notification
}

func (m *matrix) Send(msg interface{}) error {
	message := msg.(string)
	return m.sendMsg(message, m.Var2)
}

func (m *matrix) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	m.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

func (m *matrix) OnSuccess(s *types.Service) {
	if !s.Online || !s.SuccessNotified {
		m.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		var msg interface{}
		if s.UpdateNotify {
			s.UpdateNotify = false
		}
		msg = s.DownText
		m.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
	}
}

func (m *matrix) OnSave() error {
	// Validate event type
	if m.Var2 != "m.text" && m.Var2 != "m.notice" {
		return fmt.Errorf("only 'm.text' or 'm.notice' supported")
	}
	return nil
}

func (m *matrix) OnTest() error {
	return m.sendMsg("Testing Matrix Notifier on you Statping server", m.Var2)
}

func (m *matrix) sendMsg(msg, msgType string) error {
	if msgType == "" {
		msgType = "m.text"
	}

	data := map[string]string{
		"msgtype": msgType,
		"body":    msg,
	}

	// create a unique transaction id coming from this matrix client (timestamp + counter)
	timestamp := time.Now().Nanosecond()
	url := fmt.Sprintf(m.Host+matrixUrlPath, m.Var1, timestamp)

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal json: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create new request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+m.ApiSecret)
	req.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	c.Timeout = m.Delay

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("make http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	return parseMatrixError(resp.Body)
}

func parseMatrixError(reader io.Reader) error {
	type msg struct {
		Code  string `json:"ErrCode"`
		Error string `json:"error"`
	}
	code := msg{}
	err := json.NewDecoder(reader).Decode(&code)
	if err != nil {
		return fmt.Errorf("decode json: %v", err)
	}

	switch code.Code {
	case "M_UNKNOWN_TOKEN":
		return errors.New("invalid token")
	case "M_FORBIDDEN":
		// invalid room etc
		return errors.New(code.Error)
	case "M_UNRECOGNIZED":
		// matrix handled request -> homeserver url probably includes invalid subdirectory
		return errors.New("invalid homeserver url")
	default:
		return fmt.Errorf("%s: %s", code.Code, code.Error)
	}
}
