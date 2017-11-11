package main

import "fmt"

func change(n []int) {
	n[0] = 234
	fmt.Println(n)
}

func main() {
	m := make([]int, 1, 4)
	fmt.Println(m)
	change(m)
	fmt.Println(m)

}
