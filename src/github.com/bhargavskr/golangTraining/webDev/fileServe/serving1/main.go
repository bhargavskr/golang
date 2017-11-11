package main

import (
	"io"
	"net/http"
	"os"

	"log"
)

func a(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func b(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("dog.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	io.Copy(w, file)
}

func main() {

	http.HandleFunc("/", a)
	http.HandleFunc("/dog.jpg", b)

	http.ListenAndServe(":8080", nil)
}
