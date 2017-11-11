package main

import "fmt"

func main() {
	var x [256]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for i, a := range x {
		fmt.Printf("%v - %T - %b \n", a, a, a)
		if i > 50 {
			break
		}
	}
}
