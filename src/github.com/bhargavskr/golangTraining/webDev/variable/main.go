package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("./*.html")

	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "index.html", "Hi There")

}
