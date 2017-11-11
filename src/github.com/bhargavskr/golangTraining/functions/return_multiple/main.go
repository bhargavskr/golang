package main

import "fmt"

func greet(name, adverb string) (string, string) {

	return fmt.Sprint(name, adverb), fmt.Sprint(adverb, name)
}

func main() {

	fmt.Println(greet("Ram", "Good"))
	fmt.Println(greet("Hari", "Good"))
}
