package main

import "fmt"

func main() {
	var x int64

	increment := func() int64 {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

}
