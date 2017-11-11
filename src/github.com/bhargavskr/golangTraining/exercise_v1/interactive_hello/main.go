package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your Name: ")
	fmt.Scan(&name)

	fmt.Println("Hello ", name)

}
