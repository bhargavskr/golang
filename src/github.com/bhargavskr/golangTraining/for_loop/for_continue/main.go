package main

import "fmt"

func main() {
	i := 0
	for {
		i++

		if i%2 == 0 {
			continue
		}
		fmt.Print(i, "\t")

		if i >= 50 {
			break
		}
	}
}
