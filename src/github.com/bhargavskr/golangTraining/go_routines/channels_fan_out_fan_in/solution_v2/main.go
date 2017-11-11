package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)
	i := 0
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9) {
		i++
		fmt.Println(i, n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(ch))

	for _, c := range ch {
		go func(c1 <-chan int) {

			for n := range c1 {
				out <- n
			}
			wg.Done()
		}(c)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}
