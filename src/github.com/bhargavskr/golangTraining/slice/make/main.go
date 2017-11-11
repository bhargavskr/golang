package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)

	customerNumber[0] = 278
	customerNumber[1] = 28
	customerNumber[2] = 2873

	fmt.Println(customerNumber[0])

	fmt.Println(customerNumber[1])

	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)

	greeting[0] = "Good Morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	//greeting[3] = "Bon dias!"
	greeting = append(greeting, "Bon dias!")
	fmt.Println(greeting[2])

}
