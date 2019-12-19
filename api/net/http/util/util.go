package util

import (
	"fmt"
	"net/http"
)

var (
	htmlStart = []byte("<html><body>")
	htmlEnd   = []byte("</body></html>")
	homer     = []byte(`<footer><a href="/">Go back home</a></footer>`)
)

// Check will check err and print err msg
func Check(err error) {
	if err != nil {
		fmt.Println("Error occur", err.Error())
	}
}

// FormatPage formats a page with html start and end
func FormatPage(page string, isHome bool) []byte {
	p := []byte(page)

	p = append(htmlStart, p...)
	if !isHome {
		p = append(p, homer...)
	}
	p = append(p, htmlEnd...)

	return p
}

// NotFoundHandler returns a 404 page
func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	page := `<h2>404 NotFound</h2>`
	p := FormatPage(page, false)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write(p)
}

// SafeHandler returns a closure which handles panic in handler function
func SafeHandler(fnc func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	safe := func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic error:", err)
			}
		}()
		fnc(w, req)
	}
	return http.HandlerFunc(safe)
}
