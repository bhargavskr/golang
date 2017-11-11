package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for j := 0; j < n; j++ {
		go func() {

			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()

	}

	go func() {
		for j := 0; j < n; j++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
