package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {

	a := people{"kjsd", "sfd", "ewrw", "eqweqsf"}

	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)

}
