package main

import (
	"fmt"
	"strconv"
	"sync"
)

var counter int64
var mutex sync.Mutex

func main() {
	c := incrementor(2)

	for i := range c {
		counter++
		fmt.Println(i)
	}
	fmt.Println(" counter ", counter)
}

func incrementor(n int) chan string {

	c := make(chan string)
	done := make(chan bool)
	for index := 0; index < n; index++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing", k)

			}
			done <- true

		}(index)

	}

	go func() {
		for index := 0; index < n; index++ {
			<-done
		}
		close(c)
	}()
	return c
}

//IVO BALBAERT  The Way to Go
