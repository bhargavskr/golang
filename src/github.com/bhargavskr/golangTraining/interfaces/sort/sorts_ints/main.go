package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{23, 3543, 213, 2434, 66}

	fmt.Println(a)
	sort.Ints(a)
	//	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)

}
