package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func a(w http.ResponseWriter, r *http.Request) {
	data := "default Route"
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func b(w http.ResponseWriter, r *http.Request) {
	data := "My Dog"
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func c(w http.ResponseWriter, r *http.Request) {
	data := "Bhargav"
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog/", b)
	http.HandleFunc("/me/", c)
	defer http.ListenAndServe(":8080", nil)
}
