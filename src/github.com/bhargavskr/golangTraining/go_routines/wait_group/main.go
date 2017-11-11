package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {

	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}