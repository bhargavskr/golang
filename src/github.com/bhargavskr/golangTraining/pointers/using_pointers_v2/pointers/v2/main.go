package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Println("x  address in zero", &z)

}

func main() {

	x := 877
	fmt.Printf("%p x  address in main \n ", &x)
	zero(&x)
	fmt.Println("x - ", x)
	fmt.Println("x  address in main", &x)

}
