package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	myOtherSlice := []int{134, 4, 4, 4}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
}
