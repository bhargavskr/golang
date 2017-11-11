package main

import "fmt"

func main() {

	switch "Mendy" {
	case "Mendy":
		fmt.Println("Hi Mendy")
		fallthrough
	case "Danny":
		fmt.Println("Hi Danny")
		fallthrough
	case "Joe":
		fmt.Println("Hi Joe")
	default:
		fmt.Println("You are great")
	}

}
