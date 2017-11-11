package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		factorial(i)

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
		fmt.Println("factorial of n:", m)
	}
}
