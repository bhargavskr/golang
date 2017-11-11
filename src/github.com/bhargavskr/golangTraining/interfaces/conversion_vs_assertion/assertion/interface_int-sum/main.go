package main

import "fmt"

func main() {
	var name interface{} = 8

	fmt.Println(name.(int) + 8)

}
