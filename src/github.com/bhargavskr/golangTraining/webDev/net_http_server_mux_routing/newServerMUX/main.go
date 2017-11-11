package main

import (
	"io"
	"net/http"
)

type a int
type b int

func (par a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Dog")
}

func (par b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Cat")

}

func main() {
	var atype a
	var btype b
	var mux = http.NewServeMux()

	mux.Handle("/dog/", atype)
	mux.Handle("/cat", btype)

	http.ListenAndServe(":8080", mux)
}
