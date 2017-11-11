package main

import (
	"fmt"
)

func makeGreeting() func() string {
	return func() string {
		return "Welcome"
	}
}

func main() {

	greeting := makeGreeting()
	fmt.Println(greeting())
	fmt.Printf("%T \n", greeting)
}

// func expression
