package main

import (
	"fmt"
)

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	m := ""
	if n > 10 {
		m = fmt.Sprint("x is greater than 10")
	} else {
		m = fmt.Sprint("x is less than 10")

	}
	return m
}
