package main

import "fmt"

func sample(n int) (float64, bool) {

	return float64(n) / 2, n%2 == 0
}

func main() {
	n := 35
	a, b := sample(n)

	fmt.Println(n, " : ", a, " : ", b)
}
