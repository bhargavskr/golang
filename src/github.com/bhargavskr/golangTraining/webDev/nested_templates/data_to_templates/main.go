package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	//	sages := map[string]string{"India": "Gandhi", "USA": "MLK", "Love": "Jesus"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Hi")

	if err != nil {
		log.Fatalln(err)
	}
}
