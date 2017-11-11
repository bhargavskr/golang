package main

import "fmt"

func change(n *string) {
	fmt.Println(n)
	fmt.Println(*n)
	*n = "You"
	fmt.Println(n)
	fmt.Println(*n)

}

func main() {
	x := "Me"
	fmt.Println(&x)
	change(&x)
	fmt.Println(&x)
	fmt.Println(x)
}
