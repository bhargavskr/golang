package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	var myGreeting1 = map[string]string{}
	myGreeting1["Tim"] = "Good Morning"
	myGreeting1["Jenny"] = "Bonjour"

	fmt.Println(myGreeting1)

}
