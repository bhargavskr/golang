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

	sages := []string{"Gandhi", "MLK", "Jesus"}

	data := struct {
		Sages []string
		Name  string
	}{
		sages,
		"Bhargav",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)

	if err != nil {
		log.Fatalln(err)
	}
}
