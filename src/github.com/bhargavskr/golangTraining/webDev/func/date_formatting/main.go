package main

import (
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("./*.html"))
}

var fm = template.FuncMap{
	"formatTime": formatTime,
}

func formatTime(t time.Time) string {

	return t.Format("01-02-2006")
}

func main() {

	tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())

}
