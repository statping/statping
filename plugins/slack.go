package plugins

import (
	"net/http"
	"strconv"
)

const (
	SLACK_TABLE = "plugin_slack"
)

func init() {
	Add("slack")
	AddRoute("install_slack", "GET", InstallSlack)
	AddRoute("uninstall_slack", "GET", UninstallSlack)
	AddRoute("save_slack", "POST", SaveSettings)
}

type Slack struct {
	Key string
	Secret string
	Enabled bool
	Channel string
}

func InstallSlack(w http.ResponseWriter, r *http.Request) {
	CreateTable()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func UninstallSlack(w http.ResponseWriter, r *http.Request) {
	DropTable()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.PostForm.Get("key")
	secret := r.PostForm.Get("secret")
	enabled, _ := strconv.ParseBool(r.PostForm.Get("enabled"))
	channel := r.PostForm.Get("channel")

	slack := &Slack {
		key,
		secret,
		enabled,
		channel,
	}

	slack.UpdateTable()

	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}


func CreateTable() {
	sql := "CREATE TABLE "+SLACK_TABLE+" (enabled BOOLEAN, api_key text, api_secret text, channel text);"
	db.QueryRow(sql).Scan()
}

func (s *Slack) UpdateTable() {
	sql := "CREATE TABLE "+SLACK_TABLE+" (enabled BOOLEAN, api_key text, api_key text, channel text);"
	db.QueryRow(sql).Scan()
}

func DropTable() {
	sql := "DROP TABLE "+SLACK_TABLE+";"
	db.QueryRow(sql).Scan()
}
