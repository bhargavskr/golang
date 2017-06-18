package main

import "fmt"

func main() {
	for i := 60; i < 170; i++ {
		fmt.Printf("%d - %b - %x \t %q \n", i, i, i, i)

	}

}
