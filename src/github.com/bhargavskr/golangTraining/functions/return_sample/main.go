package main

import "fmt"

func greet(name, adverb string) string {
	return fmt.Sprint(name, adverb)
}

func main() {

	fmt.Println(greet("Ram", "Good"))
	fmt.Println(greet("Hari", "Good"))
}
