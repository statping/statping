package handlers

import (
	"errors"
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/configs"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func processSetupHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if core.App.Setup {
		sendErrorJson(errors.New("Statping has already been setup"), w, r)
		return
	}

	confgs, err := configs.LoadConfigForm(r)
	if err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	project := r.PostForm.Get("project")
	description := r.PostForm.Get("description")
	domain := r.PostForm.Get("domain")
	newsletter := r.PostForm.Get("newsletter")
	sendNews, _ := strconv.ParseBool(newsletter)

	log.WithFields(utils.ToFields(core.App, confgs)).Debugln("new configs posted")

	if err = configs.ConnectConfigs(confgs, false); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if err := confgs.Save(utils.Directory); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	exists := confgs.Db.HasTable("core")
	if !exists {
		if err := confgs.DropDatabase(); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := confgs.CreateDatabase(); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := configs.CreateAdminUser(confgs); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := configs.TriggerSamples(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	if err = confgs.MigrateDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	log.Infoln("Migrating Notifiers...")
	notifiers.InitNotifiers()

	c := &core.Core{
		Name:        project,
		Description: description,
		ApiSecret:   utils.Params.GetString("API_SECRET"),
		Domain:      domain,
		Version:     core.App.Version,
		Started:     utils.Now(),
		CreatedAt:   utils.Now(),
		UseCdn:      null.NewNullBool(false),
		Footer:      null.NewNullString(""),
		Language:    confgs.Language,
	}

	log.Infoln("Creating new Core")
	if err := c.Create(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	core.App = c

	if sendNews {
		log.Infof("Sending email address %s to newsletter server", confgs.Email)
		if err := registerNews(confgs.Email, confgs.Domain); err != nil {
			log.Errorln(err)
		}
	}

	log.Infoln("Initializing new Statping instance")

	if _, err := services.SelectAllServices(true); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	go services.CheckServices()

	core.App.Setup = true

	CacheStorage.Delete("/")
	resetCookies()
	time.Sleep(2 * time.Second)
	out := struct {
		Message string            `json:"message"`
		Config  *configs.DbConfig `json:"config"`
	}{
		"success",
		confgs,
	}
	returnJson(out, w, r)
}

func registerNews(email, domain string) error {
	if email == "" {
		return nil
	}
	v := url.Values{}
	v.Set("email", email)
	v.Set("domain", domain)
	v.Set("timezone", "UTC")
	rb := strings.NewReader(v.Encode())
	resp, err := http.Post("https://news.statping.com/new", "application/x-www-form-urlencoded", rb)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
