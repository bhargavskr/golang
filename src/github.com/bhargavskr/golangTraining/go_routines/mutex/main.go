package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("bar: ")
	wg.Wait()
	fmt.Println(" counter ", counter)
}

func incrementor(name string) {

	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(name, i, " counter ", counter)
		mutex.Unlock()
	}

	wg.Done()
}
