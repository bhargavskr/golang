package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []string{"kjsd", "sfd", "ewrw", "eqweqsf"}

	fmt.Println(a)
	sort.Strings(a)
	//	sort.Sort(sort.StringSlice(a))
	fmt.Println(a)

}
