package main

import "fmt"

func main() {
	n := 35
	sample := func(n int) (float64, bool) {

		return float64(n) / 2, n%2 == 0
	}

	a, b := sample(n)

	fmt.Println(n, " : ", a, " : ", b)
}
