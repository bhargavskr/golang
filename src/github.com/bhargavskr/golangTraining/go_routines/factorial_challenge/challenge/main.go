package main

import (
	"fmt"
)

func main() {
	f := factorial(5)
	fmt.Println("Factorial of 5 is ", f)
}

func factorial(n int) int {

	if n == 1 {
		return 1
	} else {
		return factorial(n-1) * n
	}

}
