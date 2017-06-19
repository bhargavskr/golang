package main

import "fmt"

const (
	_  = iota
	p  = 1 << (iota * 10)
	p1 = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%b \n ", p)
	fmt.Printf("%d \n ", p)

	fmt.Printf("%b \n ", p1)
	fmt.Printf("%d \n ", p1)
}
