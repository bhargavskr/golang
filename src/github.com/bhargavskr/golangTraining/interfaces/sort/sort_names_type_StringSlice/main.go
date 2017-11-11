package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []string{"kjsd", "sfd", "ewrw", "eqweqsf"}

	fmt.Println(a)
	sort.StringSlice(a).Sort()
	//	sort.Sort(sort.StringSlice(a))
	fmt.Println(a)

}
