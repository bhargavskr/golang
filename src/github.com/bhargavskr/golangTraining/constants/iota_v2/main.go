package main

import "fmt"

const (
	p = iota
	p1
)

func main() {

	const (
		q = iota
		q1
	)
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("p1 - ", p1)
	fmt.Println("q1 - ", q1)

}
