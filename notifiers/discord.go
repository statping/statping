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
	"fmt"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

const (
	DISCORD_METHOD = "discord"
	DISCORD_TEST   = `{"content": "This is a notification from Statup!"}`
)

type Discord struct {
	*notifier.Notification
}

var discorder = &Discord{&notifier.Notification{
	Method: DISCORD_METHOD,
	Host:   "https://discordapp.com/api/webhooks/****/*****",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Discord Webhook URL",
		Placeholder: "Insert your webhook URL here",
		DbField:     "host",
	}}},
}

// init the Discord notifier
func init() {
	err := notifier.AddNotifier(discorder)
	if err != nil {
		panic(err)
	}
}

func (u *Discord) Test() error {
	utils.Log(1, "Discord notifier loaded")
	discordPost([]byte(DISCORD_TEST))
	return nil
}

// Discord won't be using the Run() process
func (u *Discord) Run() error {
	return nil
}

// Discord won't be using the Run() process
func (u *Discord) Select() *notifier.Notification {
	return u.Notification
}

// discordPost sends an HTTP POST to the webhook URL
func discordPost(msg []byte) {
	req, _ := http.NewRequest("POST", discorder.GetValue("host"), bytes.NewBuffer(msg))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue sending Discord message to channel: %v", err))
		return
	}
	defer resp.Body.Close()
	discorder.Log(string(msg))
}

func (u *Discord) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf(`{"content": "Your service '%v' is currently failing! Reason: %v"}`, s.Name, f.Issue)
	discordPost([]byte(msg))
}

func (u *Discord) OnSuccess(s *types.Service) {

}

func (u *Discord) OnSave() error {
	msg := fmt.Sprintf(`{"content": "The Discord notifier on Statup was just updated."}`)
	discordPost([]byte(msg))
	return nil
}
