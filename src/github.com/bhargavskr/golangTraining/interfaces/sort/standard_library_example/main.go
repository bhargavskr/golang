package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d ", p.name, p.age)
}

type people []person

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func main() {

	a := []person{
		{"Bob", 22},
		{"Rick", 34},
		{"Amy", 3},
	}

	fmt.Println(a[1])
	fmt.Println(a)
	sort.Sort(people(a))
	fmt.Println(a)

}
