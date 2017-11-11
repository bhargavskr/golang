package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func a(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html chartset=UTF-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func b(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("dog.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
}
func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog.jpg", b)
	http.ListenAndServe(":8080", nil)
}
