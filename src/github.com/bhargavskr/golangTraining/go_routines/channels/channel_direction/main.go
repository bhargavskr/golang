package main

import (
	"fmt"
)

func increment() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func main() {

	input := increment()
	sum := puller(input)
	for n := range sum {
		fmt.Println(n)
	}

}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out

}
