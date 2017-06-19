package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {

	x := 877
	zero(x)
	fmt.Println("x - ", x)

}
