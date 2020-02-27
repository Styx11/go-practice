package main

import (
	"os"
	"text/template"
)

var (
	tmpl    *template.Template
	defines = `
		{{define "T0"}}This is tmpl T0, {{template "T1"}}{{end}}
		{{if "test"}}{{template "T0" printf "%s\n"}}{{end}}
		{{define "T1"}}This is tmpl T1, {{template "T2"}}{{end}}
		{{define "T2"}}This is tmpl T2, end{{end}}
	`
)

func init() {
	tmpl = template.Must(template.New("defineTest").Parse(defines))
}

func main() {
	tmpl.Execute(os.Stdout, nil)
}
