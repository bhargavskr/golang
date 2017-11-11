package main

import "fmt"

func main() {
	var y []int = []int{1, 2, 3, 56, 7}

	fmt.Println(y)

	for i, a := range y {

		fmt.Println(i, " - ", a)

	}
}
