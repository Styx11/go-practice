package main

import (
	"fmt"
	"net/http"
)

var (
	addrs     = make(map[string]bool)
	htmlStart = []byte("<html><body>")
	htmlEnd   = []byte("</body></html>")
	homer     = []byte(`<footer><a href="/">Go back home</a></footer>`)
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func formatPage(page string, isHome bool) []byte {
	p := []byte(page)

	p = append(htmlStart, p...)
	if !isHome {
		p = append(p, homer...)
	}
	p = append(p, htmlEnd...)

	return p
}

// NotFoundHandler returns a 404 page
func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	page := `<h2>404 NotFound</h2>`
	p := formatPage(page, false)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write(p)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		notFoundHandler(w, req)
		return
	}
	addr := req.RemoteAddr
	if _, ok := addrs[addr]; !ok {
		addrs[addr] = true
		fmt.Printf("Client comes in: %s\n", addr)
	}

	page := `
		<h2>Hello, This is home</h2>
		<h4>Here's what we got:</h4>
		<ul>
			<li><a href="/test1">Test1</a></li>
		</ul>
	`
	p := formatPage(page, true)

	w.Header().Set("Content-Type", "text/html")
	w.Write(p)
}

func test1Handler(w http.ResponseWriter, req *http.Request) {
	page := "<h2>Hello, World</h2>"
	p := formatPage(page, false)

	w.Header().Set("Content-Type", "text/html")
	w.Write(p)
}

// practice for example 15.10
// more detail: https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.4.md
func main() {
	// extense handler
	http.HandleFunc("/test1", test1Handler)

	// base handler
	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server at localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
