package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d - %b - %x \n", i, i, i)

	}

}
