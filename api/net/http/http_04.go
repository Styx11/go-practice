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

// practice for example 15.10
// more detail: https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.4.md
func main() {
	// base handler
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			notFoundHandler(w, req)
			return
		}
		addr := req.RemoteAddr
		if _, ok := addrs[addr]; !ok {
			addrs[addr] = true
			fmt.Printf("Client comes in: %s\n", addr)
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<h2>Hello, This is home</h2>`))
	})

	fmt.Println("Starting server at localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
