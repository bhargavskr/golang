package main

import "fmt"

func visit(numbers []int, innerFunc func(int)) {
	for _, i := range numbers {
		innerFunc(i)
	}
}

func main() {

	visit([]int{12, 23, 4, 5, 56}, func(x int) {
		fmt.Print(x, "\t")
	})

}

// functional programming
