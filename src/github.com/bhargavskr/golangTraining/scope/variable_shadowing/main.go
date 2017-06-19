package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(4)
	fmt.Println(max)
}

// shadowing variables is bad
