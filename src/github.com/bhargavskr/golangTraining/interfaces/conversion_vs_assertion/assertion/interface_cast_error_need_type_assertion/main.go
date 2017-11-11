package main

import "fmt"

func main() {
	name := 8.876

	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", int(name))

	var nam interface{} = 8

	fmt.Printf("%T \n", nam)
	//	fmt.Printf("%T \n", int(nam))
	fmt.Printf("%T \n", nam.(int))
}
