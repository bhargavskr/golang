package main

import "fmt"

func main() {

	bucket := make([]int, 1)
	fmt.Println(bucket[0])
	bucket[0] = 42
	fmt.Println(bucket[0])
	bucket[0]++
	fmt.Println(bucket[0])

}
