package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type dashboard struct {
	Services []Service
	Users    []User
	Core     *Core
}

func RunHTTPServer() {
	fmt.Println("Fusioner HTTP Server running on http://localhost:8080")
	css := http.StripPrefix("/css/", http.FileServer(cssBox.HTTPBox()))
	js := http.StripPrefix("/js/", http.FileServer(jsBox.HTTPBox()))
	http.Handle("/", http.HandlerFunc(IndexHandler))
	http.Handle("/css/", css)
	http.Handle("/js/", js)
	http.Handle("/setup", http.HandlerFunc(SetupHandler))
	http.Handle("/setup/save", http.HandlerFunc(ProcessSetupHandler))
	http.Handle("/dashboard", http.HandlerFunc(DashboardHandler))
	http.Handle("/login", http.HandlerFunc(LoginHandler))
	http.Handle("/logout", http.HandlerFunc(LogoutHandler))
	//http.Handle("/auth", http.HandlerFunc(AuthenticateHandler))
	http.Handle("/user/create", http.HandlerFunc(CreateUserHandler))
	http.Handle("/token/create", http.HandlerFunc(CreateServiceHandler))
	http.Handle("/tokens", http.HandlerFunc(ServicesHandler))
	http.Handle("/users", http.HandlerFunc(UsersHandler))
	http.ListenAndServe(":8080", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	_, auth := AuthUser(username, password)
	if auth {
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		w.WriteHeader(502)
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprintln(w, "bad")
	}
}

//func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	key := r.PostForm.Get("key")
//	secret := r.PostForm.Get("secret")
//	token := SelectToken(key, secret)
//	if token.Id != 0 {
//		go token.Hit(r)
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "plain/text")
//		fmt.Fprintln(w, token.Id)
//	} else {
//		w.WriteHeader(502)
//		w.Header().Set("Content-Type", "plain/text")
//		fmt.Fprintln(w, "bad")
//	}
//}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user := &User{
		Username: username,
		Password: password,
	}
	user.Create()
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	token := &Service{}
	token.Create()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}


func SetupHandler(w http.ResponseWriter, r *http.Request) {
	setupFile, err := tmplBox.String("setup.html")
	if err != nil {
		panic(err)
	}
	setupTmpl, err := template.New("message").Parse(setupFile)
	if err != nil {
		panic(err)
	}
	setupTmpl.Execute(w, nil)
}

type index struct {
	Services []Service
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "apizer_auth")
	//if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	if setupMode {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	dashboardFile, err := tmplBox.String("index.html")
	if err != nil {
		panic(err)
	}

	dashboardTmpl, err := template.New("message").Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
	}).Parse(dashboardFile)

	out := index{SelectAllServices()}

	dashboardTmpl.Execute(w, out)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "apizer_auth")
	//if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	dashboardFile, err := tmplBox.String("dashboard.html")
	if err != nil {
		panic(err)
	}
	dashboardTmpl, err := template.New("message").Parse(dashboardFile)
	if err != nil {
		panic(err)
	}

	out := dashboard{SelectAllServices(), SelectAllUsers(), core}

	dashboardTmpl.Execute(w, out)
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tokensFile, err := tmplBox.String("services.html")
	if err != nil {
		panic(err)
	}
	tokensTmpl, err := template.New("message").Parse(tokensFile)
	if err != nil {
		panic(err)
	}
	tokensTmpl.Execute(w, SelectAllServices())
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	usersFile, err := tmplBox.String("users.html")
	if err != nil {
		panic(err)
	}
	usersTmpl, err := template.New("message").Parse(usersFile)
	if err != nil {
		panic(err)
	}
	usersTmpl.Execute(w, SelectAllUsers())
}

func PermissionsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	permsFile, err := tmplBox.String("permissions.html")
	if err != nil {
		panic(err)
	}
	permsTmpl, err := template.New("message").Parse(permsFile)
	if err != nil {
		panic(err)
	}
	permsTmpl.Execute(w, SelectAllUsers())
}
