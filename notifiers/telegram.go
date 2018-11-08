package notifiers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
)

const (
	chatID      = -242718636
	successText = "Hallo is __statup__ success: ✅"
	errorText   = "Hallo is __statup__ error: ⛔️"
)

type telegram struct {
	*notifier.Notification
}

var telegramer = &telegram{&notifier.Notification{
	Method:      "telegram",
	Title:       "Telegram",
	Description: "Send notifications to a telegram chat.",
	Author:      "Jens Neuber",
	AuthorUrl:   "https://github.com/jensneuber",
	Delay:       time.Duration(2 * time.Second),
	ApiKey:      "https://webhooksurl.slack.com/***",
	Icon:        "fab fa-telegram",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Telegram Bot Token",
		Placeholder: "Insert your telegram bot token",
		SmallText:   "Your Telegram bot token. Create one by visiting <a href=\"https://telegram.me/BotFather\" target=\"_blank\">@BotFather</a>",
		DbField:     "api_key",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "Telegram Chat ID, Group ID, or Channel Username",
		Placeholder: "Insert your telegram chat id",
		SmallText:   "Your Telegram Chat ID, Group ID, or @channelusername. Use <a href=\"https://telegram.me/myidbot\" target=\"_blank\">@myidbot</a> on Telegram to get an ID.",
		DbField:     "var1",
		Required:    true,
	}}},
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(telegramer)
	if err != nil {
		panic(err)
	}
}

func (u *telegram) Select() *notifier.Notification {
	return u.Notification
}

func (u *telegram) Send(msg interface{}) error {
	message := msg.(string)

	utils.Log(1, fmt.Sprintf("Sending telegram message: '%v'", message))

	bot, err := tgbotapi.NewBotAPI(u.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	chatID, err := strconv.ParseInt(u.Var1, 10, 64)

	telegramMsg := tgbotapi.NewMessage(chatID, message)
	telegramMsg.ParseMode = "Markdown"
	bot.Send(telegramMsg)

	return nil
}

func (u *telegram) OnTest() error {
	bot, err := tgbotapi.NewBotAPI(u.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	chatID, err := strconv.ParseInt(u.Var1, 10, 64)

	msg := tgbotapi.NewMessage(chatID, errorText)
	msg.ParseMode = "Markdown"
	bot.Send(msg)

	return err
}

// OnFailure will trigger failing service
func (u *telegram) OnFailure(s *types.Service, f *types.Failure) {
	msg := errorText
	u.AddQueue(s.Id, msg)
	u.Online = false
}

// OnSuccess will trigger successful service
func (u *telegram) OnSuccess(s *types.Service) {
	if !u.Online {
		u.ResetUniqueQueue(s.Id)
		msg := successText
		u.AddQueue(s.Id, msg)
	}
	u.Online = true
}

// OnSave triggers when this notifier has been saved
func (u *telegram) OnSave() error {
	message := fmt.Sprintf("Notification %v is receiving updated information.", u.Method)
	u.AddQueue(0, message)
	return nil
}
