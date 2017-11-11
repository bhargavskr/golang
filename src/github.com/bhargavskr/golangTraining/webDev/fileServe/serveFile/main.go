package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html chartset=UTF-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func b(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "dog.jpg")
}
func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog.jpg", b)
	http.ListenAndServe(":8080", nil)
}
