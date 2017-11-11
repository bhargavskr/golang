package main

import "fmt"

func main() {
	x := 42

	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Inner Scope"
		fmt.Println(y)
	}
	//	fmt.Println(y)

}
