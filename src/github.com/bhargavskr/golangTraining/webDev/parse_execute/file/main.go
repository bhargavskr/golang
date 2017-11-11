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

	fi, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error in creating file ", err)
	}

	defer fi.Close()

	err = tpl.Execute(fi, nil)
	if err != nil {
		log.Fatal("Error in wrting to file ", err)
	}

}
