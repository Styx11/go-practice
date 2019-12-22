package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	nameReg     *regexp.Regexp
	htmlStart   = []byte("<html><body>")
	htmlEnd     = []byte("</body></html>")
	homer       = []byte(`<footer><a href="/">Go back home</a></footer>`)
	visitedAddr = map[string]bool{}
)

func init() {
	nameReg = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error occur:", err.Error())
	}
}

// append htmlStart and htmlEnd
func formatPage(p []byte, isHome bool) []byte {
	if !isHome {
		p = append(p, homer...)
	}
	p = append(p, htmlEnd...)
	page := append(htmlStart, p...)
	return page
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	// visiter's remote address will be print only once
	addr := req.RemoteAddr
	if _, ok := visitedAddr[addr]; !ok {
		visitedAddr[addr] = true
		fmt.Println("Client comes in:", addr)
	}

	homePage := []byte(`
		<h2>Welcome Home</h2>
		<p>Pages:</p>
		<ul>
			<li><a href="/hello">To Hello</a></li>
			<li><a href="/shouthello">To ShoutHello</a></li>
		</ul>
	`)
	w.Header().Set("Content-Type", "text/html")
	w.Write(formatPage(homePage, true))
}

// notFoundHandler handles every unregistered route
func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	notFoundPage := []byte(`
		<h2>404 Not Found</h2>
	`)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write(formatPage(notFoundPage, false))
}

// http_01 implements practice 1 in:
// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.2.md
func main() {
	fmt.Println("Starting server at localhost:8080")

	// base handlers
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if path != "/" {
			notFoundHandler(w, req) //every unregistered route will be notFound
			return
		}
		homeHandler(w, req)
	})

	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}
