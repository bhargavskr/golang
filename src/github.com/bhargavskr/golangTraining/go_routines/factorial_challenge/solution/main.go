package main

import (
	"fmt"
)

func main() {
	f := factorial(5)
	fmt.Println("Factorial of 5 is ", f)
}

func factorial(n int) int {
	c := make(chan int)
	go func() {
		for i := n; i >= 1; i-- {
			c <- i
		}
		close(c)
	}()
	f := 1
	for m := range c {
		f *= m
	}
	return f
}
