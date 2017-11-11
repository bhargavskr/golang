package main

import "fmt"

func main() {

	y := make([]int, 0, 5)

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i := 0; i < 80; i++ {
		y = append(y, i)
		fmt.Println("Length: ", len(y), " Capacity: ", cap(y), " Value ", y[i])
	}

}
