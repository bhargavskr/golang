package main

import (
	"fmt"
)

func increment(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range num {
			out <- i
		}
		close(out)
	}()
	return out
}
func main() {

	for n := range puller(increment(1, 2, 3)) {
		fmt.Println(n)
	}

}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out

}
