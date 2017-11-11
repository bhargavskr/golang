package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []string{"kjsd", "sfd", "ewrw", "eqweqsf"}

	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	//	sort.Sort(sort.StringSlice(a))
	fmt.Println(a)

	a1 := []string{"kjsd", "sfd", "ewrw", "eqweqsf"}

	sort.Sort(sort.StringSlice(a1))
	fmt.Println(a1)
	fmt.Printf("%T \n", a1)
	a1 = sort.StringSlice(a1)
	fmt.Printf("%T \n", a1)

	t := sort.StringSlice(a1)
	fmt.Printf("%T \n", t)

	fmt.Printf("%T \n", sort.StringSlice(a1))
	//fmt.Printf("%T \n", sort.Sort(sort.StringSlice(a1)))
	fmt.Printf("%T \n", sort.Reverse(sort.StringSlice(a1)))
	i := sort.Reverse(sort.StringSlice(a1))
	fmt.Println(i)

	fmt.Printf("%T \n", i)

	sort.Sort(i)
	fmt.Println(i)

}
