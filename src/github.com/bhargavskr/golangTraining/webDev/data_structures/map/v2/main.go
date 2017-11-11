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

func main() {

	sages := map[string]string{"India": "Gandhi", "USA": "MLK", "Love": "Jesus"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sages)

	if err != nil {
		log.Fatalln(err)
	}
}
