package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error in parsing file ", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error in parsing file ", err)
	}
}
