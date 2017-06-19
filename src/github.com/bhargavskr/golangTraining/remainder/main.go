package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(" remainder ", x)
	if x == 1 {
		fmt.Println(" Odd Number ")
	} else {
		fmt.Println(" Even Number ")
	}

	for i := 0; i < 100; i++ {

		if i%2 == 1 {
			fmt.Println(i, " is  Odd Number ")
		} else {
			fmt.Println(i, " is Even Number ")
		}
	}
}
