package main

import "fmt"

func main() {

	fmt.Println(hashBucket("Hello", 12))

}

func hashBucket(str string, buckets int) int {
	a := int(str[2])
	return a % 12
}
