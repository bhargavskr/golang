package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Welcome")

	}
	greeting()
	fmt.Printf("%T \n", greeting)
}

// func expression
