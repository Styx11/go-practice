package main

import (
	"fmt"
	"go-practice/api/net/http/http_07/handler"
	"net/http"
)

// http_07 implements practice 15.6
// more detail: https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.6.md
func main() {
	http.Handle("/view", handler.ViewHandler)
	http.Handle("/edit/save", handler.SaveHandler)
	http.Handle("/edit", handler.EditHandler)
	http.HandleFunc("/", handler.HomeHandler)

	fmt.Println("Starting server at localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
