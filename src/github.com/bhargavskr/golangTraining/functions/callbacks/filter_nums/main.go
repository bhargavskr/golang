package main

import (
	"fmt"
)

func visit(numbers []int, innerFunc func(int) bool) []int {
	var xs []int
	for _, i := range numbers {
		if innerFunc(i) {
			xs = append(xs, i)
		}
	}
	return xs
}

func main() {

	d := visit([]int{12, 23, 4, 5, 56}, func(x int) bool {
		return x > 10
	})

	fmt.Println(d)

}

// functional programming
