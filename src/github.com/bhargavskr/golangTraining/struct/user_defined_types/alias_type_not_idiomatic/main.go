package main

import "fmt"

type foo int

func main() {

	var h foo
	h = 87
	fmt.Printf("%T \t %v \n", h, h)

}
