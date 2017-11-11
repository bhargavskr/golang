package main

import "fmt"

func greet(name string, age int32) {
	fmt.Println("Hello ", name, " of ", age)
}

func main() {

	greet("Ram", 23)
	greet("Hari", 24)
}
