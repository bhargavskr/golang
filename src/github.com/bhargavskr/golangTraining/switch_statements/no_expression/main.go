package main

import "fmt"

func main() {
	name := "Tim"
	switch {
	case name == "Tim":
		fmt.Println("Hi Mendy")
		fallthrough
	case name == "Danny":
		fmt.Println("Hi Danny")
	case name == "Joe":
		fmt.Println("Hi Joe")
	default:
		fmt.Println("You are great")
	}

}
