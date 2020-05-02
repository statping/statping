package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/statping/statping/types/errors"
	"html/template"
	"net/http"
	"path"
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
		log.Infoln("Statping HTTP Server running on http://" + host + basePath)
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
	if ok := hasSetupEnv(); ok {
		return true
	}
	if ok := hasAPIQuery(r); ok {
		return true
	}
	if ok := hasAuthorizationHeader(r); ok {
		return true
	}
	return IsFullAuthenticated(r)
}

// IsFullAuthenticated returns true if the HTTP request is authenticated. You can set the environment variable GO_ENV=test
// to bypass the admin authenticate to the dashboard features.
func IsFullAuthenticated(r *http.Request) bool {
	if ok := hasSetupEnv(); ok {
		return true
	}
	if ok := hasAPIQuery(r); ok {
		return true
	}
	if ok := hasAuthorizationHeader(r); ok {
		return true
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

// ScopeName will show private JSON fields in the API.
// It will return "admin" if request has valid admin authentication.
func ScopeName(r *http.Request) string {
	if ok := hasAPIQuery(r); ok {
		return "admin"
	}
	if ok := hasAuthorizationHeader(r); ok {
		return "admin"
	}
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
	claim, err := getJwtToken(r)
	if err != nil {
		return false
	}
	return claim.Admin
}

// IsUser returns true if the user is registered
func IsUser(r *http.Request) bool {
	if ok := hasSetupEnv(); ok {
		return true
	}
	tk, err := getJwtToken(r)
	if err != nil {
		return false
	}
	if err := tk.Valid(); err != nil {
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

func returnJson(d interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if e, ok := d.(errors.Error); ok {
		w.WriteHeader(e.Status())
		json.NewEncoder(w).Encode(e)
		return
	}
	if e, ok := d.(error); ok {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errors.New(e.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}

// error404Handler is a HTTP handler for 404 error pages
func error404Handler(w http.ResponseWriter, r *http.Request) {
	if usingSSL {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "base.gohtml", nil, nil)
}
