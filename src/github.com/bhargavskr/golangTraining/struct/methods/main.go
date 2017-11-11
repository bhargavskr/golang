package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullname() string {
	return p.last + ", " + p.first
}

func main() {

	p1 := person{"James", "Bond", 12}
	p2 := person{"Hen", "Rick", 23}

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
