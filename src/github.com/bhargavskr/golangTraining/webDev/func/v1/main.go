package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("./*.html"))
}

type sages struct {
	Motto string
	Name  string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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
