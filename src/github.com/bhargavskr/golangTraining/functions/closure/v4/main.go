package main

import "fmt"

func wrapper() func() int64 {
	var x int64
	return func() int64 {
		x++
		return x
	}
}

func main() {

	increment := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())

}
