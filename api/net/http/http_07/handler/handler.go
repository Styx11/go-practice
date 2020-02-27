package handler

import (
	"go-practice/api/net/http/http_07/page"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var (
	wd       string
	nameReg  *regexp.Regexp
	homeTmpl *template.Template
	viewTmpl *template.Template
	editTmpl *template.Template
)

func init() {
	wd, _ = os.Getwd()
	wd += "/handler"
	nameReg = regexp.MustCompile(`^[0-9a-zA-Z_]+$`)

	homeTmpl = template.Must(template.ParseFiles(wd + "/tmpls/home.html"))
	viewTmpl = template.Must(template.ParseFiles(wd + "/tmpls/view.html"))
	editTmpl = template.Must(template.ParseFiles(wd + "/tmpls/edit.html"))
}

// ViewHandler handlers all request to /view
var ViewHandler = wikiHandler(viewHandler)

// SaveHandler handlers all request to /edit/save
var SaveHandler = wikiHandler(saveHandler)

// EditHandler handlers all request to /edit
var EditHandler = wikiHandler(editHandler)

// HomeHandler handlers all request
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	homeTmpl.Execute(w, page.Pages)
}

// WikiHandler return a HandlerFunc that checks whether name's vaild
func wikiHandler(f func(w http.ResponseWriter, req *http.Request, pname string)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		pname := req.FormValue("page")
		if !nameReg.Match([]byte(pname)) {
			http.NotFound(w, req)
			return
		}
		f(w, req, pname)
	})
}

func viewHandler(w http.ResponseWriter, req *http.Request, pageName string) {
	p := page.NewPage(pageName)
	err := p.Load()
	if err != nil {
		p.Save()
		http.Redirect(w, req, "/edit?page="+url.QueryEscape(pageName), http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	viewTmpl.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, req *http.Request, pageName string) {
	p := page.NewPage(pageName)

	body := strings.TrimSpace(req.PostFormValue("body"))
	if body == "" {
		_ = p.Del()
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}

	p.Body = []byte(body)
	p.Save()

	http.Redirect(w, req, "/view?page="+url.QueryEscape(pageName), http.StatusFound)
}

func editHandler(w http.ResponseWriter, req *http.Request, pageName string) {
	var err error
	p := page.NewPage(pageName)

	// new page
	if _, ok := page.Pages[pageName]; !ok {
		err = p.Save()
	}

	err = p.Load()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html")
	editTmpl.Execute(w, p)
}
