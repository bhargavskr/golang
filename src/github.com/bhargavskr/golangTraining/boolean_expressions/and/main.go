package main

import "fmt"

func main() {

	if true && true {
		fmt.Println("True")

	}

	if true && false {
		fmt.Println("False")
	}
}
