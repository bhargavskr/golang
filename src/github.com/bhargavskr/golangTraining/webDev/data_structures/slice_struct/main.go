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

	gandhi := sages{
		Name:  "Gandhi",
		Motto: "Peace",
	}

	mlk := sages{
		Name:  "MLK",
		Motto: "Equality",
	}

	sagesList := []sages{budda, gandhi, mlk}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sagesList)

	if err != nil {
		log.Fatalln(err)
	}
}
