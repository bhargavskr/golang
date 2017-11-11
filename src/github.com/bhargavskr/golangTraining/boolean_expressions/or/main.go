package main

import "fmt"

func main() {

	if true || false {
		fmt.Println("True")

	}

	if false || false {
		fmt.Println("False")
	}
}
