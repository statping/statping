// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
)

const (
	cookieKey = "statping_auth"
	timeout   = time.Second * 30
)

var (
	jwtKey      string
	httpServer  *http.Server
	usingSSL    bool
	mainTmpl    = `{{define "main" }} {{ template "base" . }} {{ end }}`
	templates   = []string{"base.gohtml", "head.gohtml", "nav.gohtml", "footer.gohtml", "scripts.gohtml", "form_service.gohtml", "form_notifier.gohtml", "form_integration.gohtml", "form_group.gohtml", "form_user.gohtml", "form_checkin.gohtml", "form_message.gohtml"}
	javascripts = []string{"charts.js", "chart_index.js"}
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
	return nil
}

// IsReadAuthenticated will allow Read Only authentication for some routes
func IsReadAuthenticated(r *http.Request) bool {
	if core.SetupMode {
		return false
	}
	var token string
	query := r.URL.Query()
	key := query.Get("api")
	if subtle.ConstantTimeCompare([]byte(key), []byte(core.CoreApp.ApiSecret)) == 1 {
		return true
	}
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if subtle.ConstantTimeCompare([]byte(token), []byte(core.CoreApp.ApiSecret)) == 1 {
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
	if core.CoreApp == nil {
		return true
	}
	if core.SetupMode {
		return false
	}
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if subtle.ConstantTimeCompare([]byte(token), []byte(core.CoreApp.ApiSecret)) == 1 {
			return true
		}
	}
	return IsAdmin(r)
}

func getJwtAuth(r *http.Request) (bool, string) {
	c, err := r.Cookie(cookieKey)
	if err != nil {
		utils.Log.Errorln(err)
		if err == http.ErrNoCookie {
			return false, ""
		}
		return false, ""
	}
	tknStr := c.Value
	var claims JwtClaim
	tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		utils.Log.Errorln("error getting jwt token: ", err)
		if err == jwt.ErrSignatureInvalid {
			return false, ""
		}
		return false, ""
	}
	if !tkn.Valid {
		utils.Log.Errorln("token is not valid")
		return false, ""
	}
	return claims.Admin, claims.Username
}

// IsAdmin returns true if the user session is an administrator
func IsAdmin(r *http.Request) bool {
	if core.SetupMode {
		return false
	}
	admin, username := getJwtAuth(r)
	if username == "" {
		return false
	}
	return admin
}

// IsUser returns true if the user is registered
func IsUser(r *http.Request) bool {
	if core.SetupMode {
		return false
	}
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	ff, username := getJwtAuth(r)
	fmt.Println(ff, username)
	return username != ""
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
	// render all javascript files
	for _, temp := range javascripts {
		tmp, _ := source.JsBox.String(temp)
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
	render, err := source.JsBox.String(file)
	if err != nil {
		log.Errorln(err)
	}
	if usingSSL {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
	t := template.New("charts")
	t.Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"Services": func() []types.ServiceInterface {
			return core.CoreApp.Services
		},
	})
	if _, err := t.Parse(render); err != nil {
		log.Errorln(err)
	}
	if err := t.Execute(w, data); err != nil {
		log.Errorln(err)
	}
}

func returnJson(d interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func safeTypes(obj interface{}) []string {
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		obj = &obj
	}
	switch v := obj.(type) {
	case types.Service:
		return types.SafeService
	default:
		fmt.Printf("%T\n", v)
	}
	return nil
}

func expandServices(s []types.ServiceInterface) []*types.Service {
	var services []*types.Service
	for _, v := range s {
		services = append(services, v.Select())
	}
	return services
}

func toSafeJson(input interface{}) map[string]interface{} {
	thisData := make(map[string]interface{})
	t := reflect.TypeOf(input)
	elem := reflect.ValueOf(input)

	d, _ := json.Marshal(input)

	var raw map[string]*json.RawMessage
	json.Unmarshal(d, &raw)

	if t.Kind() == reflect.Ptr {
		input = &input
	}

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("scope")
		jsonTag := field.Tag.Get("json")

		tags := strings.Split(tag, ",")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		trueValue := elem.Field(i).Interface()
		trueValue = fixValue(field, trueValue)

		if tag == "" {
			thisData[jsonTag] = trueValue
			continue
		}

		if isPublic(tags) {
			thisData[jsonTag] = trueValue
		}

		fmt.Printf("%d. %v (%v), tags: '%v'\n", i, field.Name, field.Type.Name(), tags)
	}
	return thisData
}

func returnSafeJson(w http.ResponseWriter, r *http.Request, input interface{}) {
	if reflect.ValueOf(input).Kind() == reflect.Slice {
		alldata := make([]map[string]interface{}, 0, 1)
		s := reflect.ValueOf(input)
		for i := 0; i < s.Len(); i++ {
			alldata = append(alldata, toSafeJson(s.Index(i).Interface()))
		}
		returnJson(alldata, w, r)
		return
	}
	returnJson(input, w, r)
	return
}

func fixValue(field reflect.StructField, val interface{}) interface{} {
	typeName := field.Type.Name()
	switch typeName {
	case "NullString":
		nullItem := val.(types.NullString)
		return nullItem.String
	case "NullBool":
		nullItem := val.(types.NullBool)
		return nullItem.Bool
	case "NullFloat64":
		nullItem := val.(types.NullFloat64)
		return nullItem.Float64
	case "NullInt64":
		nullItem := val.(types.NullInt64)
		return nullItem.Int64
	default:
		return val
	}
}

func isPublic(tags []string) bool {
	for _, v := range tags {
		if v == "public" {
			return true
		}
	}
	return false
}

// error404Handler is a HTTP handler for 404 error pages
func error404Handler(w http.ResponseWriter, r *http.Request) {
	if usingSSL {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	}
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "error_404.gohtml", nil, nil)
}
