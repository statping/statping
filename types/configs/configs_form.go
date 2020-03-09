package configs

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"net/http"
)

func LoadConfigForm(r *http.Request) (*DbConfig, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	dbHost := r.PostForm.Get("db_host")
	dbUser := r.PostForm.Get("db_user")
	dbPass := r.PostForm.Get("db_password")
	dbDatabase := r.PostForm.Get("db_database")
	dbConn := r.PostForm.Get("db_connection")
	dbPort := utils.ToInt(r.PostForm.Get("db_port"))
	project := r.PostForm.Get("project")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	description := r.PostForm.Get("description")
	domain := r.PostForm.Get("domain")
	email := r.PostForm.Get("email")

	if project == "" || username == "" || password == "" {
		err := errors.New("Missing required elements on setup form")
		return nil, err
	}

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
		Error:       nil,
		Location:    utils.Directory,
	}

	return confg, nil

}
