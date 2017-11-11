package main

import "fmt"

func factorial(n int) int {
	res := 1
	if n == 1 {
		res = 1
	} else {
		res = factorial(n-1) * n
	}

	return res

}

func main() {

	fmt.Println(factorial(5))

}
