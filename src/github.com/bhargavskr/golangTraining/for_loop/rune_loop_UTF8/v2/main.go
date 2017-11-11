package main

import "fmt"

func main() {
	for i := 50000; i < 50500; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	ex := 'a'
	fmt.Printf("%v \n", ex)
	fmt.Printf("%T \n", ex)
	fmt.Printf("%T \n", rune(ex))

	ex1 := "a"
	fmt.Printf("%v \n", ex1)
	fmt.Printf("%T \n", ex1)

}

// an int32 which represents a UTF 8 value
