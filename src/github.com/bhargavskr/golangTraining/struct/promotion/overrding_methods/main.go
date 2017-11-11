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

func (p person) greetings() {
	fmt.Println("Hi ", p.first)
}

func (p DoubleZero) greetings() {
	fmt.Println("Hi ", p.first)
}

func (p person) fullname() string {
	return p.last + ", " + p.first
}

func main() {

	p1 := DoubleZero{person{"James", "Bond", 12}, "Nick", "Male"}
	p2 := DoubleZero{person{"Hen", "Rick", 23}, "Steve", "Male"}

	p1.greetings()
	p1.person.greetings()
	p2.greetings()
	p2.person.greetings()
}
