package main

import "fmt"

func change(n []int) {
	fmt.Println("inner: ", n)

	n[0] = 2233
	fmt.Println("inner ", n)

}

func main() {
	x := []int{1, 2, 34}
	fmt.Println("main ", x)
	change(x)
	fmt.Println("main ", x)
}
