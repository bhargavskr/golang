package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("bskr", "this is from ravuri")
	w.Header().Set("Content-Type", "text/html; chartset=utf-8")
	fmt.Fprintln(w, "<h1> Custome Response</h1>")

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
