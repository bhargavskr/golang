package main

import (
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5)

	c1 := sq(in)
	c2 := sq(in)

	res := merge(c1, c2)

	for n := range res {
		println(n)
	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
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
