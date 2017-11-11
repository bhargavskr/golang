package main

import "fmt"

func change(n *int) {
	fmt.Println(n)
	fmt.Println(*n)
	*n = 7867
	fmt.Println(n)
	fmt.Println(*n)

}

func main() {
	x := 87
	fmt.Println(&x)
	change(&x)
	fmt.Println(&x)
	fmt.Println(x)
}
