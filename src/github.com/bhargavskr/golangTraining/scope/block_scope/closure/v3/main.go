package main

import "fmt"

func main() {

	var x = 1

	var increment = func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}
