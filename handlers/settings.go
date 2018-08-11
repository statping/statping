package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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
	core.CoreApp.UseCdn = (r.PostForm.Get("enable_cdn") == "on")
	core.CoreApp, _ = core.UpdateCore(core.CoreApp)
	core.OnSettingsSaved(core.CoreApp.ToCore())
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
	source.SaveAsset(theme, ".", "scss/base.scss")
	source.SaveAsset(variables, ".", "scss/variables.scss")
	source.CompileSASS(".")
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	dir := utils.Directory
	source.CreateAllAssets(dir)
	err := source.CompileSASS(dir)
	if err != nil {
		source.CopyToPublic(source.CssBox, dir+"/assets/css", "base.css")
		utils.Log(2, "Default 'base.css' was insert because SASS did not work.")
	}
	source.UsingAssets = true
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func DeleteAssetsHandler(w http.ResponseWriter, req *http.Request) {
	if !IsAuthenticated(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	source.DeleteAllAssets(".")
	source.UsingAssets = false
	LocalizedAssets(r)
	http.Redirect(w, req, "/settings", http.StatusSeeOther)
}

func SaveNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	r.ParseForm()

	notifierId := vars["id"]
	enabled := r.PostForm.Get("enable")

	host := r.PostForm.Get("host")
	port := int(utils.StringInt(r.PostForm.Get("port")))
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	var1 := r.PostForm.Get("var1")
	var2 := r.PostForm.Get("var2")
	apiKey := r.PostForm.Get("api_key")
	apiSecret := r.PostForm.Get("api_secret")
	limits := int(utils.StringInt(r.PostForm.Get("limits")))
	notifer := notifiers.SelectNotifier(utils.StringInt(notifierId)).Select()

	if host != "" {
		notifer.Host = host
	}
	if port != 0 {
		notifer.Port = port
	}
	if username != "" {
		notifer.Username = username
	}
	if password != "" && password != "##########" {
		notifer.Password = password
	}
	if var1 != "" {
		notifer.Var1 = var1
	}
	if var2 != "" {
		notifer.Var2 = var2
	}
	if apiKey != "" {
		notifer.ApiKey = apiKey
	}
	if apiSecret != "" {
		notifer.ApiSecret = apiSecret
	}
	if limits != 0 {
		notifer.Limits = limits
	}
	if enabled == "on" {
		notifer.Enabled = true
	} else {
		notifer.Enabled = false
	}
	notifer, err = notifer.Update()
	if err != nil {
		utils.Log(3, err)
	}

	if notifer.Enabled {
		notify := notifiers.SelectNotifier(notifer.Id)
		go notify.Run()
	}

	utils.Log(1, fmt.Sprintf("Notifier saved: %v", notifer))

	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
