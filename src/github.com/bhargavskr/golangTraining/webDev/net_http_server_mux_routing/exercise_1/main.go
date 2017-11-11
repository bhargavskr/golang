package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Default")

}

func b(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Dog")

}
func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bhargav")

}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog/", b)
	http.HandleFunc("/me/", c)
	defer http.ListenAndServe(":8080", nil)
}
