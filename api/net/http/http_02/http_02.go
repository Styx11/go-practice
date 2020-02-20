package main

import (
	"fmt"
	"net/http"
)

// an empty struct that implement Handler interface
type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

// http_01 implements practice 2 in:
// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.2.md
func main() {
	var h hello
	fmt.Println("Starting server at localhost:99")

	http.Handle("/hello", h)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/hello", http.StatusSeeOther)
	})
	http.ListenAndServe(":99", nil)
}
