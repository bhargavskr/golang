package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//	tpl, err := template.ParseGlob("templates/*.gmao")
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal("Error in parsing file ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatal("Error  ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatal("Error  ", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "three.gmao", nil)
	if err != nil {
		log.Fatal("Error  ", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error  ", err)
	}

}
