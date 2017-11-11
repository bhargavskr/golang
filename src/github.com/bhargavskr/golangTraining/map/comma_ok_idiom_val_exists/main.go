package main

import "fmt"

func main() {

	greeting := map[int]string{

		0: "Good Morning",
		1: "Bonjour!",
		2: "dias!",
		3: "Bon dias!",
	}

	fmt.Println(greeting)

	if val, exists := greeting[2]; exists {
		delete(greeting, 2)
		fmt.Println("val = ", val)
		fmt.Println("exists =", exists)

	} else {
		fmt.Println("The value does not exists")
		fmt.Println("val = ", val)
		fmt.Println("exists =", exists)

	}
	fmt.Println(greeting)

}
