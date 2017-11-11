package main

import "fmt"

var x int64

func increment() int64 {
	x++
	return x
}

func main() {

	fmt.Println(increment())
	fmt.Println(increment())

}
