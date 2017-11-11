package main

import (
	"fmt"
)

func main() {
	for l := 0; l < 2; l++ {
		for i := 1; i <= 65; i++ {
			factorial(i)

		}
	}
}

func factorial(n int) {
	c := make(chan int)

	go func() {
		total := 1
		for i := n; i >= 1; i-- {
			total *= i

		}
		c <- total
		close(c)
	}()
	for m := range c {
		fmt.Println("factorial of ", n, " : ", m)
	}
}
