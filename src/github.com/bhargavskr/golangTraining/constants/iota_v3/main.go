package main

import "fmt"

const (
	_  = iota
	p  = iota * 10
	p1 = iota * 10
)

func main() {

	const (
		_  = iota
		q  = iota * 10
		q1 = iota * 10
	)
	fmt.Println("p - ", p)
	fmt.Println("p1 - ", p1)
	fmt.Println("q - ", q)

	fmt.Println("q1 - ", q1)

}
