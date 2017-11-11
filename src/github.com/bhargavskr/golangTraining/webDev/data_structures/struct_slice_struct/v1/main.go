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

type cars struct {
	Model string
	Make  string
	Doors int
}

type items struct {
	SagesList []sages
	Transport []cars
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

	ford := cars{
		Model: "F150",
		Make:  "Ford",
		Doors: 2,
	}

	c := cars{
		Model: "Corolla",
		Make:  "Toyota",
		Doors: 2,
	}

	item := items{
		SagesList: []sages{budda, gandhi, mlk},

		Transport: []cars{ford, c},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", item)

	if err != nil {
		log.Fatalln(err)
	}
}
