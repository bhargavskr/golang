package main

import "fmt"

func change(n string) {
	fmt.Println(n)

	n = "world"
	fmt.Println(n)

}

func main() {
	x := "hello"
	fmt.Println(&x)
	change(x)
	fmt.Println(&x)
	fmt.Println(x)
}
