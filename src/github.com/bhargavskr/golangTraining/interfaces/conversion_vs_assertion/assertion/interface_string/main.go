package main

import "fmt"

func main() {
	var name interface{} = "Syndeny"
	val, ok := name.(string)
	if ok {
		fmt.Printf("%T \n", val)
	} else {
		fmt.Println("Not is a integer")
	}

}
