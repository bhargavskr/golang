package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64
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

		atomic.AddInt64(&counter, 1)
		fmt.Println(name, i, " counter ", counter)

	}

	wg.Done()
}

//IVO BALBAERT  The Way to Go
