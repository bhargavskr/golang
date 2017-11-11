package main

import "fmt"

func greet(name, adverb string) (s string) {
	s = fmt.Sprint(name, adverb)
	return
}

func main() {

	fmt.Println(greet("Ram", "Good"))
	fmt.Println(greet("Hari", "Good"))
}
