package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html chartset=UTF-8")
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", a)
	http.ListenAndServe(":8080", nil)
}
