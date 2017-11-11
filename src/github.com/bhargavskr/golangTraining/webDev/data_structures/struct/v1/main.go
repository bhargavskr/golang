package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))
}

type sages struct {
	Motto string
	Name  string
}

func main() {

	//	sages := map[string]string{"India": "Gandhi", "USA": "MLK", "Love": "Jesus"}

	budda := sages{
		Name:  "Buddha",
		Motto: "Love people",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", budda)

	if err != nil {
		log.Fatalln(err)
	}
}
