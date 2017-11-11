package main

import (
	"fmt"
)

func increment(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}
func main() {

	input := increment("Foo: ")
	input1 := increment("Bar: ")

	sum := puller(input)
	sum1 := puller(input1)
	fmt.Println("Sum is ", <-sum+<-sum1)
	for n := range sum {
		fmt.Println(n)
	}

}

func puller(c chan int) chan int {
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
