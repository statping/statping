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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"strings"
	"time"
)

type discord struct {
	*notifier.Notification
}

var Discorder = &discord{&notifier.Notification{
	Method:      "discord",
	Title:       "discord",
	Description: "Send notifications to your discord channel using discord webhooks. Insert your discord channel Webhook URL to receive notifications. Based on the <a href=\"https://discordapp.com/developers/docs/resources/Webhook\">discord webhooker API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(5 * time.Second),
	Host:        "https://discordapp.com/api/webhooks/****/*****",
	Icon:        "fab fa-discord",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "discord webhooker URL",
		Placeholder: "Insert your Webhook URL here",
		DbField:     "host",
	}}},
}

// Send will send a HTTP Post to the discord API. It accepts type: []byte
func (u *discord) Send(msg interface{}) error {
	message := msg.(string)
	_, _, err := utils.HttpRequest(Discorder.GetValue("host"), "POST", "application/json", nil, strings.NewReader(message), time.Duration(10*time.Second), true)
	return err
}

func (u *discord) Select() *notifier.Notification {
	return u.Notification
}

// OnFailure will trigger failing service
func (u *discord) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf(`{"content": "Your service '%v' is currently failing! Reason: %v"}`, s.Name, f.Issue)
	u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnSuccess will trigger successful service
func (u *discord) OnSuccess(s *types.Service) {
	if !s.Online || !s.SuccessNotified {
		u.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		var msg interface{}
		if s.UpdateNotify {
			s.UpdateNotify = false
		}
		msg = s.DownText

		u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
	}
}

// OnSave triggers when this notifier has been saved
func (u *discord) OnSave() error {
	msg := fmt.Sprintf(`{"content": "The discord notifier on Statping was just updated."}`)
	u.AddQueue("saved", msg)
	return nil
}

// OnSave triggers when this notifier has been saved
func (u *discord) OnTest() error {
	outError := errors.New("Incorrect discord URL, please confirm URL is correct")
	message := `{"content": "Testing the discord notifier"}`
	contents, _, err := utils.HttpRequest(Discorder.Host, "POST", "application/json", nil, bytes.NewBuffer([]byte(message)), time.Duration(10*time.Second), true)
	if string(contents) == "" {
		return nil
	}
	var d discordTestJson
	err = json.Unmarshal(contents, &d)
	if err != nil {
		return outError
	}
	if d.Code == 0 {
		return outError
	}
	fmt.Println("discord: ", string(contents))
	return nil
}

type discordTestJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
