package configs

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"net/http"
	"strconv"
)

func LoadConfigForm(r *http.Request) (*DbConfig, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	g := r.PostForm.Get
	dbHost := g("db_host")
	dbUser := g("db_user")
	dbPass := g("db_password")
	dbDatabase := g("db_database")
	dbConn := g("db_connection")
	dbPort := utils.ToInt(g("db_port"))
	project := g("project")
	username := g("username")
	password := g("password")
	description := g("description")
	domain := g("domain")
	email := g("email")
	language := g("language")
	reports, _ := strconv.ParseBool(g("send_reports"))

	if project == "" || username == "" || password == "" {
		err := errors.New("Missing required elements on setup form")
		return nil, err
	}

	p := utils.Params
	p.Set("DB_CONN", dbConn)
	p.Set("DB_HOST", dbHost)
	p.Set("DB_USER", dbUser)
	p.Set("DB_PORT", dbPort)
	p.Set("DB_PASS", dbPass)
	p.Set("DB_DATABASE", dbDatabase)
	p.Set("NAME", project)
	p.Set("DESCRIPTION", description)
	p.Set("LANGUAGE", language)
	p.Set("ALLOW_REPORTS", reports)
	p.Set("ADMIN_USER", username)
	p.Set("ADMIN_PASSWORD", password)
	p.Set("ADMIN_EMAIL", email)

	confg := &DbConfig{
		DbConn:      dbConn,
		DbHost:      dbHost,
		DbUser:      dbUser,
		DbPass:      dbPass,
		DbData:      dbDatabase,
		DbPort:      int(dbPort),
		Project:     project,
		Description: description,
		Domain:      domain,
		Username:    username,
		Password:    password,
		Email:       email,
		Location:    utils.Directory,
		Language:    language,
		SendReports: reports,
	}

	return confg, nil
}
