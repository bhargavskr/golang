package main

import (
	"io"
	"net/http"
)

func dogPic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<!--not serving from our server-->
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
		`)
}

func main() {
	http.HandleFunc("/dog/", dogPic)
	http.ListenAndServe(":8080", nil)
}
