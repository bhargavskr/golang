package main

import "fmt"

func sample(n int) (int, bool) {

	return n / 2, n%2 == 0
}

func main() {
	n := 34
	a, b := sample(n)

	fmt.Println(n, " : ", a, " : ", b)
}
