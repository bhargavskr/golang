package main

import "fmt"

func main() {

	var myGreeting map[string]string
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

}
