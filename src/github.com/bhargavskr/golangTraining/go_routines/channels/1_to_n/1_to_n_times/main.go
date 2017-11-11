package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {

		for i := 0; i < 50; i++ {
			c <- i
		}
		close(c)
	}()

	for j := 0; j < n; j++ {
		go func() {

			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()

	}

	for k := 0; k < n; k++ {
		<-done

	}
}
