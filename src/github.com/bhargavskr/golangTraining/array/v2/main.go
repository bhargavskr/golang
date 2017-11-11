package main

import "fmt"

func main() {
	var y [45]int
	fmt.Println(y)
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)

	fmt.Println(x[42])
}
