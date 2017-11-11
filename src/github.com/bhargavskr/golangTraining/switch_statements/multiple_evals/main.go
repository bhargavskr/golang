package main

import "fmt"

func main() {

	switch "Tim" {
	case "Mendy", "Tim":
		fmt.Println("Hi Mendy")
		fallthrough
	case "Danny":
		fmt.Println("Hi Danny")
	case "Joe":
		fmt.Println("Hi Joe")
	default:
		fmt.Println("You are great")
	}

}
