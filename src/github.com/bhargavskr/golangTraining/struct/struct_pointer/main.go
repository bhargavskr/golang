package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 66}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.age)
	fmt.Println(p1.name)

}
