package main

import "fmt"

func main() {

	var num, deno int32
	fmt.Print("Enter a numerator: ")
	fmt.Scan(&num)
	fmt.Print("Enter a denominator: ")
	fmt.Scan(&deno)

	fmt.Println(num, "/", deno, " : ", num/deno)
}
