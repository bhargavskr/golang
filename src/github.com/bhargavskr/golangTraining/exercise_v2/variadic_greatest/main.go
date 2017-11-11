package main

import "fmt"

func sample(n ...int) int {
	var max int

	for _, g := range n {
		if g > max {
			max = g
		}
	}

	return max
}

func main() {
	slic := []int{1, 24, 564, 324, 42, 32}
	a := sample(slic...)

	fmt.Println(slic, " :> ", a)
}
