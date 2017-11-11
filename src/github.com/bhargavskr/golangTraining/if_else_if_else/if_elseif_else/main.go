package main

import "fmt"

func main() {

	b := false
	if food := "Biryani"; b {
		fmt.Println("True ", food)
	} else if false {
		fmt.Println("False ", food)

	} else {
		fmt.Println("Last Statement")

	}
	//	fmt.Println("True ", food)
}
