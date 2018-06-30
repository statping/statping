package handlers

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func PluginsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
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
	fmt.Println(core.CoreApp.Communications)

	ExecuteResponse(w, r, "plugins.html", core.CoreApp)
}

func SaveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
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
	auth := IsAuthenticated(r)
	if !auth {
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
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	core.CreateAllAssets()
	core.UsingAssets = true
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveEmailSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	emailer := core.SelectCommunication(1)

	r.ParseForm()
	emailer.Host = r.PostForm.Get("host")
	emailer.Username = r.PostForm.Get("username")
	emailer.Password = r.PostForm.Get("password")
	emailer.Port = int(utils.StringInt(r.PostForm.Get("port")))
	emailer.Var1 = r.PostForm.Get("address")
	core.Update(emailer)

	//sample := &Email{
	//	To:       SessionUser(r).Email,
	//	Subject:  "Sample Email",
	//	Template: "error.html",
	//}
	//AddEmail(sample)

	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
