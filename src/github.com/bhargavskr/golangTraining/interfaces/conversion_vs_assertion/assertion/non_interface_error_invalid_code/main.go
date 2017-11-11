package main

import "fmt"

func main() {
	name := "Syndeny"
	val, ok := name.(string)
	if ok {
		fmt.Printf("%q \n", val)
	} else {
		fmt.Println("Not is a integer")
	}

}
