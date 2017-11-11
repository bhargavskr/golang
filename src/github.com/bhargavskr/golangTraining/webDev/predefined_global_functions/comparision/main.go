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

	//	sages := map[string]string{"India": "Gandhi", "USA": "MLK", "Love": "Jesus"}

	data := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)

	if err != nil {
		log.Fatalln(err)
	}
}
