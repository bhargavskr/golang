package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	person
	first  string
	gender string
}

func (p person) fullname() string {
	return p.last + ", " + p.first
}

func main() {

	p1 := DoubleZero{person{"James", "Bond", 12}, "Nick", "Male"}
	p2 := DoubleZero{person{"Hen", "Rick", 23}, "Steve", "Male"}

	fmt.Println(p1.fullname(), p1.first, p1.last, p1.gender, p1.person.first)
	fmt.Println(p2.fullname(), p2.first, p2.last, p2.gender, p2.person.first)
}
