package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html chartset=UTF-8")
	io.WriteString(w, `<img src="/resources/dog.jpg">`)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", a)
	http.ListenAndServe(":8080", nil)
}
