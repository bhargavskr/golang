package main

import "fmt"

func main() {
	str := "World"
	a := rune(str[3])
	fmt.Println(a)

	fmt.Printf("%T \n", a)

}
