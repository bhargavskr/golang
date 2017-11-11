package main

import (
	"io"
	"net/http"
)

func atype(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Dog")
}

func btype(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Cat")

}

func main() {

	http.HandleFunc("/dog/", atype)
	http.HandleFunc("/cat", btype)

	http.ListenAndServe(":8080", nil)
}
