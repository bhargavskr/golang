package main

import "fmt"

func change(n map[string]int) {
	n["me"] = 234
	fmt.Println(n)
}

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	change(m)
	fmt.Println(m)

}
