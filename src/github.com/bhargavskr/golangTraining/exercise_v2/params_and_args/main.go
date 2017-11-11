package main

import "fmt"

func sample(n ...int) {
	fmt.Println()
	for _, i := range n {
		fmt.Print(i, " \t ")
	}
}

func main() {
	sample(1, 2)
	sample(1, 23, 2)
	aSlice := []int{12, 4, 55, 552}
	sample(aSlice...)
	sample()
}
