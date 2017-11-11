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

	http.Handle("/dog/", http.HandlerFunc(atype))
	http.Handle("/cat", http.HandlerFunc(btype))

	http.ListenAndServe(":8080", nil)
}
