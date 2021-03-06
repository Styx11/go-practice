package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Login means a login's info
type Login struct {
	Name     string
	PassWord string
	Remember string
}

// User means a user's info
type User struct {
	Name     string
	PassWord string
}

var (
	wd                  string
	currentUser         *User
	rootTmpl            *template.Template
	managerPublicTmpl   *template.Template
	managerOverviewTmpl *template.Template
)

func init() {
	wd, _ = os.Getwd()
	wd += "/handler"
	currentUser = &User{}

	rootTmpl = template.Must(template.ParseFiles(wd + "/templates/root.html"))
	managerPublicTmpl = template.Must(template.ParseGlob(wd + "/templates/manager/public/*.html"))
	managerOverviewTmpl = template.Must(managerPublicTmpl.ParseFiles(wd + "/templates/manager/overview.html"))
}

// RootHandler handles all request to /
var RootHandler = makeHandler(rootHandler)

// LoginHandler handles login request
var LoginHandler = makeHandler(loginHandler)

func makeHandler(f func(w http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalf("%v\n", err)
				w.WriteHeader(http.StatusBadRequest)
			}
		}()
		f(w, req)
	})
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	// 若用户已登陆，应用用户cookie
	userCookie, err := req.Cookie("user")
	if err == nil {
		userValue := strings.Split(userCookie.Value, " ")
		name := userValue[0]
		pw := userValue[1]
		http.Redirect(w, req, "/home?user="+name+"&password="+pw, http.StatusSeeOther)
		return
	}

	// 应用登陆cookie
	login := &Login{}
	cookie, err := req.Cookie("login")
	if err == nil {
		loginValue := strings.Split(cookie.Value, " ")
		login.Name = loginValue[0]
		login.PassWord = loginValue[1]
		login.Remember = loginValue[2]
	}

	w.Header().Set("Content-Type", "text/html")
	rootTmpl.Execute(w, login)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	name := req.PostFormValue("name")
	pw := req.PostFormValue("password")
	remember := req.PostFormValue("remember")
	signout := req.FormValue("signout")

	// 设置用户Cookie
	userCookie := &http.Cookie{}
	userCookie.Name = "user"
	userCookie.Value = strings.Join([]string{name, pw}, " ")
	userCookie.Path = "/"

	// 退出登陆，删除cookie
	if signout != "" {
		userCookie.MaxAge = -1
		http.SetCookie(w, userCookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	userCookie.MaxAge = 60 * 60

	// 设置登陆cookie
	loginCookie := &http.Cookie{}
	loginCookie.Name = "login"
	loginCookie.Value = strings.Join([]string{name, pw, remember}, " ")
	loginCookie.Path = "/"
	if remember != "" {
		loginCookie.MaxAge = 60 * 60
	} else {
		loginCookie.MaxAge = -1
	}

	http.SetCookie(w, loginCookie)
	http.SetCookie(w, userCookie)

	http.Redirect(w, req, "/home?user="+name+"&password="+pw, http.StatusFound)
}
