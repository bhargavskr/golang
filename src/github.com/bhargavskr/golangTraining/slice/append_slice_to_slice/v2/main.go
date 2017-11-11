package main

import "fmt"

func main() {
	mySlice := []string{"Mon", "Tue"}
	myOtherSlice := []string{"Wed", "Thur", "Sat", "Fri", "Sun"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
}
