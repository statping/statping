package plugins

import (
	"net/http"
)

const (
	SLACK_TABLE   = "plugin_slack"
	SLACK_INSTALL = "CREATE TABLE " + SLACK_TABLE + " (enabled BOOLEAN, api_key text, api_secret text, channel text);"
)

type Slack struct {
	Key           string
	Secret        string
	Enabled       bool
	Channel       string
	InstallFunc   func()
	UninstallFunc func()
	SaveFunc      func(*http.Request)
}

func init() {

	plugin := &Plugin{
		"slack",
		SLACK_INSTALL,
		InstallSlack,
		UninstallSlack,
		SaveSlack,
	}

	plugin.Add()

}

func InstallSlack() {
	CreateTable()
}

func UninstallSlack() {
	DropTable()
}

func SaveSlack() {
	//key := r.PostForm.Get("key")
	//secret := r.PostForm.Get("secret")
	//enabled, _ := strconv.ParseBool(r.PostForm.Get("enabled"))
	//channel := r.PostForm.Get("channel")

	//slack.UpdateTable()

}

func CreateTable() {
	sql := "CREATE TABLE " + SLACK_TABLE + " (enabled BOOLEAN, api_key text, api_secret text, channel text);"
	db.QueryRow(sql).Scan()
}

func (s *Slack) UpdateTable() {
	sql := "CREATE TABLE " + SLACK_TABLE + " (enabled BOOLEAN, api_key text, api_key text, channel text);"
	db.QueryRow(sql).Scan()
}

func DropTable() {
	sql := "DROP TABLE " + SLACK_TABLE + ";"
	db.QueryRow(sql).Scan()
}
