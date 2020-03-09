// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"crypto/subtle"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/statping/statping/types/core"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/statping/statping/source"
	"github.com/statping/statping/utils"
)

const (
	cookieKey = "statping_auth"
	timeout   = time.Second * 30
)

var (
	jwtKey     string
	httpServer *http.Server
	usingSSL   bool
	mainTmpl   = `{{define "main" }} {{ template "base" . }} {{ end }}`
	templates  = []string{"base.gohtml"}
)

// RunHTTPServer will start a HTTP server on a specific IP and port
func RunHTTPServer(ip string, port int) error {
	host := fmt.Sprintf("%v:%v", ip, port)

	key := utils.FileExists(utils.Directory + "/server.key")
	cert := utils.FileExists(utils.Directory + "/server.crt")

	if key && cert {
		log.Infoln("server.cert and server.key was found in root directory! Starting in SSL mode.")
		log.Infoln(fmt.Sprintf("Statping Secure HTTPS Server running on https://%v:%v", ip, 443))
		usingSSL = true
	} else {
		log.Infoln("Statping HTTP Server running on http://" + host)
	}

	router = Router()
	resetCookies()

	if usingSSL {
		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		srv := &http.Server{
			Addr:         fmt.Sprintf("%v:%v", ip, 443),
			Handler:      router,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			WriteTimeout: timeout,
			ReadTimeout:  timeout,
			IdleTimeout:  timeout,
		}
		return srv.ListenAndServeTLS(utils.Directory+"/server.crt", utils.Directory+"/server.key")
	} else {
		httpServer = &http.Server{
			Addr:         host,
			WriteTimeout: timeout,
			ReadTimeout:  timeout,
			IdleTimeout:  timeout,
			Handler:      router,
		}
		httpServer.SetKeepAlivesEnabled(false)
		return httpServer.ListenAndServe()
	}
}

// IsReadAuthenticated will allow Read Only authentication for some routes
func IsReadAuthenticated(r *http.Request) bool {
	if !core.App.Setup {
		return false
	}
	var token string
	query := r.URL.Query()
	key := query.Get("api")
	if subtle.ConstantTimeCompare([]byte(key), []byte(core.App.ApiSecret)) == 1 {
		return true
	}
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if subtle.ConstantTimeCompare([]byte(token), []byte(core.App.ApiSecret)) == 1 {
			return true
		}
	}
	return IsFullAuthenticated(r)
}

// IsFullAuthenticated returns true if the HTTP request is authenticated. You can set the environment variable GO_ENV=test
// to bypass the admin authenticate to the dashboard features.
func IsFullAuthenticated(r *http.Request) bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if core.App == nil {
		return true
	}
	if !core.App.Setup {
		return false
	}
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if subtle.ConstantTimeCompare([]byte(token), []byte(core.App.ApiSecret)) == 1 {
			return true
		}
	}
	return IsAdmin(r)
}

func getJwtToken(r *http.Request) (JwtClaim, error) {
	c, err := r.Cookie(cookieKey)
	if err != nil {
		if err == http.ErrNoCookie {
			return JwtClaim{}, err
		}
		return JwtClaim{}, err
	}
	tknStr := c.Value
	var claims JwtClaim
	tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return JwtClaim{}, err
		}
		return JwtClaim{}, err
	}
	if !tkn.Valid {
		return claims, errors.New("token is not valid")
	}
	return claims, err
}

func ScopeName(r *http.Request) string {
	claim, err := getJwtToken(r)
	if err != nil {
		return ""
	}
	if claim.Admin {
		return "admin"
	}
	return "user"
}

// IsAdmin returns true if the user session is an administrator
func IsAdmin(r *http.Request) bool {
	if !core.App.Setup {
		return false
	}
	if utils.Getenv("GO_ENV", false).(bool) {
		return true
	}
	claim, err := getJwtToken(r)
	if err != nil {
		return false
	}
	return claim.Admin
}

// IsUser returns true if the user is registered
func IsUser(r *http.Request) bool {
	if !core.App.Setup {
		return false
	}
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	_, err := getJwtToken(r)
	if err != nil {
		return false
	}
	return true
}

func loadTemplate(w http.ResponseWriter, r *http.Request) (*template.Template, error) {
	var err error
	mainTemplate := template.New("main")
	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	mainTemplate.Funcs(handlerFuncs(w, r))
	// render all templates
	for _, temp := range templates {
		tmp, _ := source.TmplBox.String(temp)
		mainTemplate, err = mainTemplate.Parse(tmp)
		if err != nil {
			log.Errorln(err)
			return nil, err
		}
	}
	return mainTemplate, err
}

// ExecuteResponse will render a HTTP response for the front end user
func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}, redirect interface{}) {
	if url, ok := redirect.(string); ok {
		http.Redirect(w, r, path.Join(basePath, url), http.StatusSeeOther)
		return
	}
	if usingSSL {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
	mainTemplate, err := loadTemplate(w, r)
	if err != nil {
		log.Errorln(err)
	}
	render, err := source.TmplBox.String(file)
	if err != nil {
		log.Errorln(err)
	}
	// render the page requested
	if _, err := mainTemplate.Parse(render); err != nil {
		log.Errorln(err)
	}
	// execute the template
	if err := mainTemplate.Execute(w, data); err != nil {
		log.Errorln(err)
	}
}

// executeJSResponse will render a Javascript response
func executeJSResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	//render, err := source.JsBox.String(file)
	//if err != nil {
	//	log.Errorln(err)
	//}
	//if usingSSL {
	//	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//}
	//t := template.New("charts")
	//t.Funcs(template.FuncMap{
	//	"safe": func(html string) template.HTML {
	//		return template.HTML(html)
	//	},
	//	"Services": func() []services.ServiceInterface {
	//		return core.CoreApp.Services
	//	},
	//})
	//if _, err := t.Parse(render); err != nil {
	//	log.Errorln(err)
	//}
	//if err := t.Execute(w, data); err != nil {
	//	log.Errorln(err)
	//}
}

func returnJson(d interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

// error404Handler is a HTTP handler for 404 error pages
func error404Handler(w http.ResponseWriter, r *http.Request) {
	if usingSSL {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "index.html", nil, nil)
}
