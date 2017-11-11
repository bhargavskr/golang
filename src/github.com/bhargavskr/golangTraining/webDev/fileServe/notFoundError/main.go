package main

import (
	"fmt"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "Check terminal")
}
func main() {

	http.HandleFunc("/", a)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
