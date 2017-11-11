package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	home := Page{
		Title:   "New Page",
		Heading: "Hi",
		Input:   `<script>alert("Hi")</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
