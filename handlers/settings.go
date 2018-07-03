package handlers

import (
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/notifications"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func PluginsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//CoreApp.FetchPluginRepo()

	//var pluginFields []PluginSelect
	//
	//for _, p := range allPlugins {
	//	fields := structs.Map(p.GetInfo())
	//
	//	pluginFields = append(pluginFields, PluginSelect{p.GetInfo().Name, p.GetForm(), fields})
	//}

	//CoreApp.PluginFields = pluginFields

	ExecuteResponse(w, r, "settings.html", core.CoreApp)
}

func SaveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	name := r.PostForm.Get("project")
	if name != "" {
		core.CoreApp.Name = name
	}
	description := r.PostForm.Get("description")
	if description != core.CoreApp.Description {
		core.CoreApp.Description = description
	}
	style := r.PostForm.Get("style")
	if style != core.CoreApp.Style {
		core.CoreApp.Style = style
	}
	footer := r.PostForm.Get("footer")
	if footer != core.CoreApp.Footer {
		core.CoreApp.Footer = footer
	}
	domain := r.PostForm.Get("domain")
	if domain != core.CoreApp.Domain {
		core.CoreApp.Domain = domain
	}
	core.CoreApp.Update()
	core.OnSettingsSaved(core.CoreApp)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveSASSHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	theme := r.PostForm.Get("theme")
	variables := r.PostForm.Get("variables")
	core.SaveAsset(theme, "scss/base.scss")
	core.SaveAsset(variables, "scss/variables.scss")
	core.CompileSASS()
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	core.CreateAllAssets()
	core.UsingAssets = true
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveEmailSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	emailer := core.SelectCommunication(1)
	r.ParseForm()
	smtpHost := r.PostForm.Get("host")
	smtpUser := r.PostForm.Get("username")
	smtpPass := r.PostForm.Get("password")
	smtpPort := int(utils.StringInt(r.PostForm.Get("port")))
	smtpOutgoing := r.PostForm.Get("address")
	enabled := r.PostForm.Get("enable_email")

	emailer.Host = smtpHost
	emailer.Username = smtpUser
	if smtpPass != "#######################" {
		emailer.Password = smtpPass
	}
	emailer.Port = smtpPort
	emailer.Var1 = smtpOutgoing
	emailer.Enabled = false
	if enabled == "on" {
		emailer.Enabled = true
	}
	core.Update(emailer)

	sample := &types.Email{
		To:       SessionUser(r).Email,
		Subject:  "Test Email",
		Template: "message.html",
		From:     emailer.Var1,
	}
	notifications.LoadEmailer(emailer)
	notifications.SendEmail(core.EmailBox, sample)
	notifications.EmailComm = emailer
	if emailer.Enabled {
		utils.Log(1, "Starting Email Routine, 1 unique email per 60 seconds")
		go notifications.EmailRoutine()
	}

	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveSlackSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	slack := core.SelectCommunication(2)
	r.ParseForm()
	slackWebhook := r.PostForm.Get("slack_url")
	enable := r.PostForm.Get("enable_slack")
	slack.Enabled = false
	if enable == "on" && slackWebhook != "" {
		slack.Enabled = true
		go notifications.SlackRoutine()
	}
	slack.Host = slackWebhook
	core.Update(slack)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
