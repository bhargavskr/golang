package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good Morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	greeting[3] = "Bon dias!"

	fmt.Println(greeting[6])

}
